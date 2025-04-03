import { FC, useState } from "react";
import { Todo } from "../types/todo";
import Styles from "../styles/todoList.module.scss";
import { Check, XIcon } from "lucide-react";
import clsx from "clsx";
import Modal from "react-modal";

export const TodoList: FC<{
  todos: Todo[];
  toggleCompleted: (id: number) => void;
  deleteTodo: (id: number) => void;
}> = ({ todos, toggleCompleted, deleteTodo }) => {
  return (
    <div className={Styles.list}>
      <div className={Styles.marking}>
        <div className={Styles.completed}>completed</div>
        <div className={Styles.title}>title</div>
        <div className={Styles.priority}>priority</div>
        <div className={Styles.createdAt}>created at</div>
        <div className={Styles.timeToComplete}>time to complete</div>
      </div>
      {todos.map((todo) => (
        <TodoItem
          key={todo.id}
          deleteTodo={deleteTodo}
          toggleCompleted={toggleCompleted}
          todo={todo}
        />
      ))}
    </div>
  );
};

const timeToComplete = (todo: Todo) => {
  if (!todo.timeToComplete) {
    return "unlim";
  }

  const createDate = new Date(todo.createDate);
  const deadline = new Date(
    createDate.getTime() + todo.timeToComplete * 60 * 60 * 1000
  );
  const remainingTime = deadline.getTime() - Date.now();

  const hours = Math.floor(Math.abs(remainingTime) / (1000 * 60 * 60));
  const minutes = Math.floor(
    (Math.abs(remainingTime) % (1000 * 60 * 60)) / (1000 * 60)
  );

  return remainingTime >= 0
    ? `${hours}h ${minutes}m`
    : `overdue by: ${hours}h ${minutes}m`;
};

const convertTime = (time: string) => {
  const createDate = new Date(time);

  if (isNaN(createDate.getTime())) {
    console.error(`invalid date format: ${time}`);
    return "invalid date";
  }

  const hours = createDate.getHours().toString().padStart(2, "0");
  const minutes = createDate.getMinutes().toString().padStart(2, "0");
  const day = createDate.getDate().toString().padStart(2, "0");
  const month = (createDate.getMonth() + 1).toString().padStart(2, "0"); // месяцы с 0
  const year = createDate.getFullYear().toString().slice(-2); // последние 2 цифры года

  return `${day}.${month}.${year} ${hours}:${minutes}`;
};

const TodoItem: FC<{
  todo: Todo;
  toggleCompleted: (id: number) => void;
  deleteTodo: (id: number) => void;
}> = ({ todo, toggleCompleted, deleteTodo }) => {
  const [deleteModal, setDeleteModal] = useState(false);

  return (
    <div className={Styles.item}>
      <div className={Styles.content}>
        <div className={Styles.completed}>
          <div
            onClick={() => toggleCompleted(todo.id)}
            className={clsx(
              Styles.completedBox,
              todo.completed && Styles.completedBoxActive
            )}
          >
            {todo.completed && <Check size={20} />}
          </div>
        </div>
        <div
          className={clsx(Styles.title, todo.completed && Styles.titleActive)}
        >
          {todo.title}
        </div>
        <div className={Styles.priority}>
          {todo.priority ? todo.priority : "none"}
        </div>
        <div className={Styles.createdAt}> {convertTime(todo.createDate)}</div>
        <div className={Styles.timeToComplete}>
          {todo.completed ? "completed" : timeToComplete(todo)}
        </div>
      </div>

      <div onClick={() => setDeleteModal(true)} className={Styles.deleteBut}>
        <XIcon />
      </div>

      <Modal
        style={{
          overlay: { backgroundColor: "rgba(0, 0, 0, 0.5)" },
          content: {
            padding: "20px",
            width: "400px",
            background: "#404040",
            height: "200px",
            margin: "auto",
          },
        }}
        isOpen={deleteModal}
        onRequestClose={() => setDeleteModal(false)}
        contentLabel="123"
      >
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            justifyContent: "space-between",
            height: "100%",
          }}
        >
          <h2> do you want to delete {todo.title} ? </h2>
          <div
            style={{
              display: "flex",
              justifyContent: "flex-end",
              padding: "0  20px",
            }}
          >
            <button style={{ padding: "4px", margin: "10px", fontSize: '20px', cursor: 'pointer' }} 
              onClick={() => {
                deleteTodo(todo.id)
              }}
            > yes </button>
            <button
            onClick={() => {
              setDeleteModal(false)
            }}
            style={{ padding: "4px", margin: "10px", fontSize: '20px', cursor: 'pointer' }}> cancel </button>
          </div>
        </div>
      </Modal>
    </div>
  );
};
