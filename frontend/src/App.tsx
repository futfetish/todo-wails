import './App.scss'
import {GetTodos , AddTodo} from '../wailsjs/go/main/App'
import { useEffect } from 'react';

function App() {
  useEffect(() => {
    GetTodos(null).then((data) => console.log(111,data) ).catch(console.error);
  }, []);

  return (
    <>
      <div>hewwo world</div>
    </>
  )
}

export default App
