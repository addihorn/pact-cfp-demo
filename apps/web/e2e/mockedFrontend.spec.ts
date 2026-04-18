import {test, expect } from '@playwright/test'
import {Pact, SpecificationVersion} from '@pact-foundation/pact'
import path from 'path'
import { boolean, eachLike, integer, string, datetime } from '@pact-foundation/pact/src/v3/matchers'
import { getRandomValues, randomBytes, randomInt } from 'crypto'




test.describe("The UI", () => {
   const provider = new Pact({
        dir: path.resolve(process.cwd(), 'pacts'),
        consumer: 'ToDoService-Frontend',
        provider: 'ToDoService-Backend',
        spec: SpecificationVersion.SPECIFICATION_VERSION_V4,
    })
 

    test("can load existing todos", async ({page, request}) =>  {
        await page.route('**/api/v1/todos', async (route) => {
            provider.setup()
            const stub = provider.addInteraction()
                .testName('reading the todo list generated a valid list in the ui')
                .given('at least three todos are in the database')
                .uponReceiving('a request to read all todos')
                .withRequest('GET', '/api/v1/todos', (builder) => {
                    builder.headers(<any>route.request().headers())
                })
                .willRespondWith(200, (builder) => {
                        builder.headers({'Content-Type':'application/json'})
                        builder.jsonBody({data: 
                            eachLike(
                                {
                                    id: integer(),
                                    title: string("New ToDo"),
                                    description: string("This a new description of a todo"),
                                    completed: boolean(false),
                                    createdAt: datetime("yyyy-MM-dd'T'HH:mm:ss'Z'", '2026-07-21T18:17:43Z'),
                                    updatedAt: datetime("yyyy-MM-dd'T'HH:mm:ss'Z'", '2026-07-22T00:17:43Z'),
                                }, 3)
                        })
                })            
            await stub.executeTest( async (mockserver) => {
                const response = await request.get(`${mockserver.url}/api/v1/todos`, {headers: <any>route.request().headers()})
                const json = await response.json()
                await route.fulfill({response, json})
            })
            
        })
        await page.goto('')
        await expect(page.locator('xpath=//section[@class="todos-panel"]').getByRole('list').getByRole('listitem')).toHaveCount(3)
    })
})
