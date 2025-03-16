import { FC, useEffect, useState } from "react";
import { DeleteTodo, GetTodos, ToggleTodo } from "../../wailsjs/go/main/App";
import { Todo } from "../types/todo";
import { TodoList } from "../components/todoList";

export const Completed: FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    const fetchTodos = async () => {
      const data = (await GetTodos(true)) as Todo[];
      console.log(data);
      if (data) {
        setTodos(data);
      }
    };

    fetchTodos();
  }, []);

  const toggleCompleted = (id: number) => {
    ToggleTodo(id);
    setTodos((prevTodos) => prevTodos.filter((todo) => todo.id !== id));
  };

  const deleteTodo = (id: number) => {
    DeleteTodo(id);
    setTodos((prevTodos) => prevTodos.filter((todo) => todo.id !== id));
  };

  return (
    <div>
      <h1>completed todos</h1>
      <TodoList
        deleteTodo={deleteTodo}
        toggleCompleted={toggleCompleted}
        todos={todos}
      />
    </div>
  );
};
