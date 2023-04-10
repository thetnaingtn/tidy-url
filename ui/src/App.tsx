import React from "react";
import IconSend from "./icons/IconSend";
import IconCopy from "./icons/IconCopy";

function App() {
  return (
    <div className="w-screen h-screen bg-[#ccebe3] flex flex-col justify-center items-center">
      <div className="flex">
        <input
          type="text"
          className="outline-none bg-slate-50 border-solid border-[#a5c3c2] border rounded-md w-96 h-9 text-ellipsis"
        />
        <button type="button" className="ml-[-30px]">
          <IconSend />
        </button>
      </div>
      <div className="flex">
        <p>test</p>
        <button type="button">
          <IconCopy />
        </button>
      </div>
    </div>
  );
}

export default App;
