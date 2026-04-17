import { FormEvent, useState } from "react";

type TodoFormProps = {
  onSubmit: (title: string, description: string) => Promise<void>;
};

export default function TodoForm({ onSubmit }: TodoFormProps) {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [busy, setBusy] = useState(false);

  const handleSubmit = async (event: FormEvent) => {
    event.preventDefault();
    if (!title.trim()) {
      return;
    }
    setBusy(true);
    try {
      await onSubmit(title.trim(), description.trim());
      setTitle("");
      setDescription("");
    } finally {
      setBusy(false);
    }
  };

  return (
    <form className="todo-form" onSubmit={handleSubmit}>
      <h2>Create Todo</h2>
      <label>
        Title
        <input value={title} onChange={(e) => setTitle(e.target.value)} placeholder="Add architecture diagram" required />
      </label>
      <label>
        Description
        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          placeholder="Optional details"
          rows={3}
        />
      </label>
      <button type="submit" disabled={busy}>
        {busy ? "Saving..." : "Add Todo"}
      </button>
    </form>
  );
}