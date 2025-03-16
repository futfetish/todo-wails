import { FC, useEffect, useState } from "react";
import Styles from "../styles/home.module.scss";
import { useForm } from "react-hook-form";
import { priorityValues, Todo } from "../types/todo";
import { TodoList } from "../components/todoList";
import {
  AddTodo,
  DeleteTodo,
  GetTodos,
  ToggleTodo,
} from "../../wailsjs/go/main/App";

export const Home: FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchTodos = async () => {
      setIsLoading(true)
      const data = (await GetTodos(null)) as Todo[];
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

  const addTodoToList = (todo: Todo) => {
    setTodos((todos) => [...todos, todo]);
  };

  const toggleCompleted = (id: number) => {
    ToggleTodo(id);
    setTodos((prevTodos) =>
      prevTodos.map((todo) =>
        todo.id === id ? { ...todo, completed: !todo.completed } : todo
      )
    );
  };

  const deleteTodo = (id: number) => {
    DeleteTodo(id);
    setTodos((prevTodos) => prevTodos.filter((todo) => todo.id !== id));
  };

  return (
    <div className={Styles.content}>
      <h1>Todo List</h1>
      <ToDoForm addTodoToList={addTodoToList} />
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

interface FormData {
  title: Todo["title"];
  timeToComplete: string;
  priority: NonNullable<Todo["priority"]> | "";
}

const ToDoForm: FC<{ addTodoToList: (todos: Todo) => void }> = ({
  addTodoToList,
}) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<FormData>();

  const onSubmit = (data: FormData) => {
    console.log(data);
    reset();

    const timeToComplete =
      data.timeToComplete.trim() === "" || isNaN(parseInt(data.timeToComplete))
        ? null
        : parseInt(data.timeToComplete);

    const priority = data.priority === "" ? null : data.priority;

    AddTodo(data.title.trim(), priority, timeToComplete).then((data) =>
      addTodoToList(data as Todo)
    );
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className={Styles.form}>
      <div className={Styles.block}>
        <label>title</label>
        <input
          {...register("title", {
            required: "title is required",
            validate: (value) =>
              value.trim() !== "" || "title cannot be empty or spaces only",
          })}
        />
        {errors.title && <p className={Styles.error}>{errors.title.message}</p>}
      </div>

      <div className={Styles.block}>
        <label>time to complete in hours (optional)</label>
        <input type="number" {...register("timeToComplete")} />
        {errors.timeToComplete && (
          <p className={Styles.error}>{errors.timeToComplete.message}</p>
        )}
      </div>

      <div className={Styles.block}>
        <label>priority (optional)</label>
        <select {...register("priority")}>
          {priorityValues.map((status) => (
            <option key={status} value={status === null ? "" : status}>
              {status === null ? "not selected" : status}
            </option>
          ))}
        </select>
        {errors.priority && (
          <p className={Styles.error}>{errors.priority.message}</p>
        )}
      </div>

      <button className={Styles.but} type="submit">
        add
      </button>
    </form>
  );
};
