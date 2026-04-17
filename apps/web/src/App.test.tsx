import { describe, expect, it, vi } from "vitest";
import { render, screen } from "@testing-library/react";

import App from "./App";



vi.mock("./api/client", () => ({
  listTodos: vi.fn(async () => []),
  createTodo: vi.fn(async () => ({})),
  patchTodo: vi.fn(async () => ({})),
  deleteTodo: vi.fn(async () => undefined)
}));

describe("App", () => {
  it("renders core sections", async () => {
    render(<App />);
    expect(await screen.findByText("Todo Control Panel")).toBeInTheDocument();
    expect(screen.getByText("Create Todo")).toBeInTheDocument();
    expect(screen.getByText("Todos")).toBeInTheDocument();
  });
});