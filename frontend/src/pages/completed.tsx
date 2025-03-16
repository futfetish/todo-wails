import { FC, useEffect, useState } from "react";
import { DeleteTodo, GetTodos, ToggleTodo } from "../../wailsjs/go/main/App";
import { Todo } from "../types/todo";
import { TodoList } from "../components/todoList";

export const Completed: FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchTodos = async () => {
      setIsLoading(true);
      const data = (await GetTodos(true)) as Todo[];
      if (!data) {
        console.error("ошибка: данные пустые");
        return;
      }
      console.log(data);
      setTodos(data);
      setIsLoading(false);
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
      {isLoading && todos.length == 0 ? (
        <h2> loading... </h2>
      ) : (
        <TodoList
          deleteTodo={deleteTodo}
          toggleCompleted={toggleCompleted}
          todos={todos}
        />
      )}
    </div>
  );
};
