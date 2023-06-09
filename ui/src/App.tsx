import { MouseEventHandler, useRef, useState } from "react";
import useSWRMutation from "swr/mutation";

import IconSend from "./icons/IconSend";
import IconCopy from "./icons/IconCopy";
import fetcher from "./fetcher";
import GithubCorner from "./components/GithubCorner/GithubCorner";
import { validateURL } from "./util";

function App() {
  const { data, trigger } = useSWRMutation(
    `${import.meta.env.VITE_API_ENDPOINT}/api/tidy`,
    fetcher
  );
  const [longUrl, setLongUrl] = useState("");
  const result = useRef<HTMLInputElement | null>(null);

  const disableButton = !longUrl || !validateURL(longUrl);

  const handleTidyUp: MouseEventHandler<HTMLButtonElement> = () => {
    trigger({
      long_url: longUrl,
    });
  };

  return (
    <>
      <GithubCorner />
      <div className="w-screen h-screen bg-[#ccebe3] flex flex-col justify-center items-center gap-3">
        <div className="flex relative w-full px-4 sm:w-min sm:px-0">
          <input
            type="text"
            value={longUrl}
            onChange={(e) => {
              setLongUrl(e.target.value);
            }}
            placeholder="Paste your long URL here..."
            className="outline-none bg-slate-50 border-solid border-[#a5c3c2] border rounded-md w-96 h-11 text-ellipsis py-2 pl-2 pr-7 text-[#a5c3c2]"
          />
          <button
            type="button"
            className="absolute right-5 sm:right-1 top-[0.6875rem] disabled:cursor-not-allowed"
            disabled={disableButton}
            onClick={handleTidyUp}
          >
            <IconSend />
          </button>
        </div>
        {data && (
          <div className="flex">
            <input
              readOnly
              ref={result}
              className="text-[#74a09e] bg-transparent outline-none w-96"
              value={data.short_url}
            />

            <button
              onClick={() => {
                result.current?.select();
                document.execCommand("copy");
              }}
              type="button"
            >
              <IconCopy />
            </button>
          </div>
        )}
      </div>
    </>
  );
}

export default App;
