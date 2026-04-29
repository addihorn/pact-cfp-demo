import {test, expect, Page, APIRequestContext} from '@playwright/test'
import {Pact, SpecificationVersion} from '@pact-foundation/pact'
import path from 'path'
import { boolean, eachLike, integer, string, datetime, like } from '@pact-foundation/pact/src/v3/matchers'




test.describe("The UI", () => {
   const provider = new Pact({
        dir: path.resolve(process.cwd(), 'pacts'),
        consumer: 'ToDoService-Frontend',
        provider: 'pactflow-bidi-provider',
        spec: SpecificationVersion.SPECIFICATION_VERSION_V4,
    })
 
    test("can load existing todos", async ({page, request}) =>  {
        await openPage(page, request)
        await expect(page.locator('xpath=//section[@class="todos-panel"]').getByRole('list').getByRole('listitem')).toHaveCount(3)
    })

    test("can patch existing todos", async ({page, request})=> {
        await openPage(page, request)
        await markAsDone(page, request)

    })

    const todoList = {data:
        eachLike(
            {
                id: integer(),
                title: string("New ToDo"),
                description: string("This a new description of a todo"),
                completed: boolean(false),
                createdAt: datetime("yyyy-MM-dd'T'HH:mm:ss'Z'", '2026-07-21T18:17:43Z'),
                updatedAt: datetime("yyyy-MM-dd'T'HH:mm:ss'Z'", '2026-07-22T00:17:43Z'),
            }, 3)
    }


    async function openPage(page: Page, request: APIRequestContext) {
            await page.route('**/api/v1/todos', async (route) => {
                provider.setup()
                const stub = provider.addInteraction()
                    .testName('reading the todo list generated a valid list in the ui')
                    .given('at least three todos are in the database')
                    .given('a todo with id 1 exists in the database')
                    .given('todo 1 is not completed ')
                    .uponReceiving('a request to read all todos')
                    .withRequest('GET', '/api/v1/todos', (builder) => {
                        builder.headers(<any>route.request().headers())
                    })
                    .willRespondWith(200, (builder) => {
                            builder.headers({'Content-Type':'application/json'})
                            builder.jsonBody(todoList.data)
                    })            
                await stub.executeTest( async (mockserver) => {
                    const response = await request.get(`${mockserver.url}/api/v1/todos`, {headers: <any>route.request().headers()})
                    const json = await response.json()
                    await route.fulfill({response, json})
                })                
            })
            await page.goto('')
            
    }

    async function markAsDone(page: Page, request: APIRequestContext) {
            await page.route('**/api/v1/todos/*', async (route) => {
                provider.setup()
                const stub = provider.addInteraction()
                    .testName('marking a Todo to done')
                    .given('a todo with id 1 exists in the database')
                    .given('todo 1 is not completed ')
                    .uponReceiving('an update for the first todo')
                    .withRequest('PATCH', '/api/v1/todos/1', (builder) => {
                        builder.headers(<any>route.request().headers())
                        builder.jsonBody(route.request().postData())
                    })
                    .willRespondWith(200, (builder) => {
                            builder.headers({'Content-Type':'application/json'})
                            builder.jsonBody({data: 
                                like({
                                        id: integer(1),
                                        title: string("New ToDo"),
                                        description: string("This a new description of a todo"),
                                        completed: boolean(true),
                                        createdAt: datetime("yyyy-MM-dd'T'HH:mm:ss'Z'", '2026-07-21T18:17:43Z'),
                                        updatedAt: datetime("yyyy-MM-dd'T'HH:mm:ss'Z'", '2026-07-22T00:17:43Z'),
                                    })
                            })
                    })            
                await stub.executeTest( async (mockserver) => {
                    const response = await request.patch(`${mockserver.url}/api/v1/todos/1`, {headers: <any>route.request().headers()})
                    const json = await response.json()
                    await route.fulfill({response, json})
                })                
            })

            await page.getByRole('button', { name: 'Mark Done' }).first().click()
    }


})


