export interface Todo {
    id: number;
    title: string;
    completed: boolean;
    dueDate?: string; // может быть null
    priority: "low" | "medium" | "high"; // соответствует enum в Go
  }