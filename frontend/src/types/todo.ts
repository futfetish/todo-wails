export const priorityValues = [null, "low", "medium", "high"] as const;

export interface Todo {
  id: number;
  title: string;
  completed: boolean;
  createDate: string;
  timeToComplete: number | null;
  priority: (typeof priorityValues)[number];
}
