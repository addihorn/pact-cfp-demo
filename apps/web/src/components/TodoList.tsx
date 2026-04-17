import { Todo } from "../types/todo";

type TodoListProps = {
  todos: Todo[];
  onToggle: (todo: Todo) => Promise<void>;
  onDelete: (id: number) => Promise<void>;
};

export default function TodoList({ todos, onToggle, onDelete }: TodoListProps) {
  if (todos.length === 0) {
    return <p className="empty-state">No todos found.</p>;
  }

  return (
    <ul className="todo-list">
      {todos.map((todo) => (
        <li key={todo.id} className={todo.completed ? "done" : ""}>
          <div>
            <strong>{todo.title}</strong>
            {todo.description ? <p>{todo.description}</p> : null}
          </div>
          <div className="todo-actions">
            <button type="button" onClick={() => onToggle(todo)}>
              {todo.completed ? "Mark Active" : "Mark Done"}
            </button>
            <button type="button" onClick={() => onDelete(todo.id)}>
              Delete
            </button>
          </div>
        </li>
      ))}
    </ul>
  );
}