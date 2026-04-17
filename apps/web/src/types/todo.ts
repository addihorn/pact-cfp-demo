export type Todo = {
  id: number;
  title: string;
  description: string;
  completed: boolean;
  createdAt: string;
  updatedAt: string;
};

export type TodoPayload = {
  title: string;
  description: string;
};

export type TodoPatchPayload = Partial<Pick<Todo, "title" | "description" | "completed">>;