import { useCallback, useEffect, useState } from "react";

import { createTodo, deleteTodo, listTodos, patchTodo } from "./api/client";
import FilterBar from "./components/FilterBar";
import TodoForm from "./components/TodoForm";
import TodoList from "./components/TodoList";
import { Todo } from "./types/todo";

type Filter = "all" | "active" | "completed";

export default function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [filter, setFilter] = useState<Filter>("all");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const loadTodos = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await listTodos(filter);
      setTodos(data);
    } catch (e) {
      setError((e as Error).message);
    } finally {
      setLoading(false);
    }
  }, [filter]);

  useEffect(() => {
    void loadTodos();
  }, [loadTodos]);

  const handleCreate = async (title: string, description: string) => {
    await createTodo({ title, description });
    await loadTodos();
  };

  const handleToggle = async (todo: Todo) => {
    const data = await patchTodo(todo.id, { completed: !todo.completed });
    //await loadTodos();
  };

  const handleDelete = async (id: number) => {
    await deleteTodo(id);
    await loadTodos();
  };

  return (
    <main className="layout">
      <header>
        <p className="tag">3-tier-system architecture demo</p>
        <h1>Todo Control Panel</h1>
      </header>

      <TodoForm onSubmit={handleCreate} />

      <section className="todos-panel">
        <div className="panel-header">
          <h2>Todos</h2>
          <FilterBar value={filter} onChange={setFilter} />
        </div>

        {loading ? <p>Loading todos...</p> : null}
        {error ? <p className="error">{error}</p> : null}
        {!loading && !error ? <TodoList todos={todos} onToggle={handleToggle} onDelete={handleDelete} /> : null}
      </section>
    </main>
  );
}