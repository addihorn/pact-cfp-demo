import { Todo, TodoPatchPayload, TodoPayload } from "../types/todo";

var API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080";

type DataResponse<T> = { data: T };

type ErrorResponse = { error: string };

export function updateBaseURL(newUrl: string): void {
  API_BASE_URL = newUrl
}

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    headers: {
      "Content-Type": "application/json",
      ...(init?.headers ?? {})
    },
    ...init
  });

  if (!response.ok) {
    let message = `Request failed with status ${response.status}`;
    try {
      const body = (await response.json()) as ErrorResponse;
      if (body.error) {
        message = body.error;
      }
    } catch {
      // ignore decode issues
    }
    throw new Error(message);
  }

  if (response.status === 204) {
    return undefined as T;
  }

  const payload = (await response.json()) as DataResponse<T>;
  return payload.data;
}

export function listTodos(filter?: "all" | "active" | "completed"): Promise<Todo[]> {
  const params = new URLSearchParams();
  if (filter === "active") {
    params.set("completed", "false");
  }
  if (filter === "completed") {
    params.set("completed", "true");
  }
  const qs = params.toString();
  return request<Todo[]>(`/api/v1/todos${qs ? `?${qs}` : ""}`);
}

export function createTodo(payload: TodoPayload): Promise<Todo> {
  return request<Todo>("/api/v1/todos", {
    method: "POST",
    body: JSON.stringify(payload)
  });
}

export function patchTodo(id: number, payload: TodoPatchPayload): Promise<Todo> {
  return request<Todo>(`/api/v1/todos/${id}`, {
    method: "PATCH",
    body: JSON.stringify(payload)
  });
}

export function deleteTodo(id: number): Promise<void> {
  return request<void>(`/api/v1/todos/${id}`, { method: "DELETE" });
}