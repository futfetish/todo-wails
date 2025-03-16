export const priorityValues = ["low", "medium", "high"] as const;

export interface Todo {
  id: number;
  title: string;
  completed: boolean;
  createDate: string;
  timeToComplete: number | null; // может быть null
  priority: (typeof priorityValues)[number];
}
