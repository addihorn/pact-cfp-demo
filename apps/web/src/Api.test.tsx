<<<<<<< Updated upstream
import { describe, expect, it } from "vitest";
import { Pact, SpecificationVersion} from '@pact-foundation/pact'
import path from 'path'
import { Todo } from "./types/todo";
import { boolean, eachLike, integer, like, string, } from '@pact-foundation/pact/src/v3/matchers'
import { datetime } from '@pact-foundation/pact/src/v3/matchers'
import { createTodo, listTodos, updateBaseURL } from "./api/client";







describe("The API", () => {

    const provider = new Pact({
        dir: path.resolve(process.cwd(), 'pacts'),
        consumer: 'ToDoService-Frontend',
        provider: 'ToDoService-Backend',
        spec: SpecificationVersion.SPECIFICATION_VERSION_V4
    })

  it("can add new ToDo's with the correct data", async () => {
    
    const stub = provider.addInteraction()
        .uponReceiving('a new todo')
        .withRequest('POST', '/api/v1/todos', (builder) => {
            builder.headers({'content-type' : 'application/json'})
            builder.jsonBody(like({title: string("new ToDo"), description: string("This is a new todo")}))
        })
        .willRespondWith(201, (builder) => {
            builder.headers({'content-type' : 'application/json'})
            builder.jsonBody(like({data: {
                    id: integer(1),
                    title: string('new ToDo'),
                    description: string('This is a new todo'),
                    completed: boolean(false),
                    createdAt: datetime("yyyy-MM-dd'T'HH:mm:ss'Z'", '2026-07-21T18:17:43Z'),
                    updatedAt: datetime("yyyy-MM-dd'T'HH:mm:ss'Z'", '2026-07-22T00:17:43Z'),
            }}))
        })
    
    const expectedTodo: Todo = {
        id:1,
        title: "new ToDo",
        description: "This is a new todo",
        completed: false,
        createdAt: "2026-07-21T18:17:43Z",
        updatedAt: "2026-07-22T00:17:43Z"
    }
   
    const newTodo = await stub.executeTest(async (mockserver) => {
        vi.stubEnv("VITE_API_BASE_URL", mockserver.url)    
       
        updateBaseURL(mockserver.url)
        return await createTodo({title: "new ToDo", description: "This is a new todo"})            
    })
    expect(newTodo).toEqual(expectedTodo)
  });

  
  it("can load a list of ToDo's", async () => {
    
    const stub = provider.addInteraction()
        .given('the system already knows at least 3 todos')
        .uponReceiving('a request to read the todo list')
        .withRequest('GET', '/api/v1/todos')
        .willRespondWith(200, (builder) => {
            builder.jsonBody({data: eachLike({   
                    id: integer(1), 
                    title: string('1'), 
                    description: string('Beschreibung'), 
                    completed: boolean(), 
                    createdAt: datetime("yyyy-MM-dd'T'HH:mm:ss'Z'", '2026-07-21T18:17:43Z'),
                    updatedAt: datetime("yyyy-MM-dd'T'HH:mm:ss'Z'", '2026-07-22T00:17:43Z'),
                }, 3)
            })
        })
    const expectedTodo: Todo = {   
        id: 1, 
        title: '1', 
        description: 'Beschreibung', 
        completed: true, 
        createdAt: '2026-07-21T18:17:43Z',
        updatedAt: '2026-07-22T00:17:43Z',
    }
    
    
    const todoList = await stub.executeTest( async (mockserver) => {        
        updateBaseURL(mockserver.url)
        return await listTodos()
        
    })
    console.log(todoList)
    expect(todoList).toContainEqual(expectedTodo)
  })
=======
import { describe, expect, it } from "vitest";
import { Pact, SpecificationVersion} from '@pact-foundation/pact'
import path from 'path'
import { Todo } from "./types/todo";
import { boolean, eachLike, integer, like, string, } from '@pact-foundation/pact/src/v3/matchers'
import { datetime } from '@pact-foundation/pact/src/v3/matchers'
import { createTodo, listTodos, updateBaseURL } from "./api/client";







describe("The API", () => {

    const provider = new Pact({
        dir: path.resolve(process.cwd(), 'pacts'),
        consumer: 'ToDoService-Frontend',
        provider: 'ToDoService-Backend',
        spec: SpecificationVersion.SPECIFICATION_VERSION_V4
    })

  it("can add new ToDo's with the correct data", async () => {
    
    const stub = provider.addInteraction()
        .uponReceiving('a new todo')
        .withRequest('POST', '/api/v1/todos', (builder) => {
            builder.headers({'content-type' : 'application/json'})
            builder.jsonBody(like({title: string("new ToDo"), description: string("This is a new todo")}))
        })
        .willRespondWith(201, (builder) => {
            builder.headers({'content-type' : 'application/json'})
            builder.jsonBody(like({data: {
                    id: integer(1),
                    title: string('new ToDo'),
                    description: string('This is a new todo'),
                    completed: boolean(false),
                    createdAt: datetime('yyyy-MM-ddThh:mm:ssZ', '2026-07-21T18:17:43Z'),
                    updatedAt: datetime('yyyy-MM-ddThh:mm:ssZ', '2026-07-22T00:17:43Z'),
            }}))
        })
    
    const expectedTodo: Todo = {
        id:1,
        title: "new ToDo",
        description: "This is a new todo",
        completed: false,
        createdAt: "2026-07-21T18:17:43Z",
        updatedAt: "2026-07-22T00:17:43Z"
    }
   
    const newTodo = await stub.executeTest(async (mockserver) => {
        vi.stubEnv("VITE_API_BASE_URL", mockserver.url)    
       
        updateBaseURL(mockserver.url)
        return await createTodo({title: "new ToDo", description: "This is a new todo"})            
    })
    expect(newTodo).toEqual(expectedTodo)
  });

  
  it("can load a list of ToDo's", async () => {
    
    const stub = provider.addInteraction()
        .given('the system already knows at least 3 todos')
        .uponReceiving('a request to read the todo list')
        .withRequest('GET', '/api/v1/todos')
        .willRespondWith(200, (builder) => {
            builder.jsonBody({data: eachLike({   
                    id: integer(1), 
                    title: string('1'), 
                    description: string('Beschreibung'), 
                    completed: boolean(), 
                    createdAt: datetime('yyyy-MM-ddThh:mm:ssZ', '2026-07-21T18:17:43Z'),
                    updatedAt: datetime('yyyy-MM-ddThh:mm:ssZ', '2026-07-22T00:17:43Z'),
                }, 3)
            })
        })
    const expectedTodo: Todo = {   
        id: 1, 
        title: '1', 
        description: 'Beschreibung', 
        completed: true, 
        createdAt: '2026-07-21T18:17:43Z',
        updatedAt: '2026-07-22T00:17:43Z',
    }
    
    
    const todoList = await stub.executeTest( async (mockserver) => {        
        updateBaseURL(mockserver.url)
        return await listTodos()
        
    })
    console.log(todoList)
    expect(todoList).toContainEqual(expectedTodo)
  })
>>>>>>> Stashed changes
});