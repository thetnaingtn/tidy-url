import React, { MouseEventHandler, useState } from "react";
import useSWRMutation from "swr/mutation";

import IconSend from "./icons/IconSend";
import IconCopy from "./icons/IconCopy";
import fetcher from "./fetcher";

function App() {
  const { data, trigger } = useSWRMutation(
    `${import.meta.env.VITE_API_ENDPOINT}/tidy`,
    fetcher
  );
  const [longUrl, setLongUrl] = useState("");

  const handleTidyUp: MouseEventHandler<HTMLButtonElement> = () => {
    trigger({
      long_url: longUrl,
    });
  };

  return (
    <div className="w-screen h-screen bg-[#ccebe3] flex flex-col justify-center items-center">
      <div className="flex">
        <input
          type="text"
          value={longUrl}
          onChange={(e) => {
            setLongUrl(e.target.value);
          }}
          placeholder="Paste Your Long URL Here..."
          className="outline-none bg-slate-50 border-solid border-[#a5c3c2] border rounded-md w-96 h-9 text-ellipsis py-2 pl-2 pr-7"
        />
        <button type="button" className="ml-[-30px]" onClick={handleTidyUp}>
          <IconSend />
        </button>
      </div>
      {data && (
        <div className="flex">
          <p>{data.short_url}</p>
          <button type="button">
            <IconCopy />
          </button>
        </div>
      )}
    </div>
  );
}

export default App;
