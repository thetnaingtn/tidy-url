import { type MouseEventHandler, useState } from "react";
import {GithubCorner} from "./components/github-corner";
import { tidyUrlServiceClient } from "./grpcweb";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { HoverCard, HoverCardContent, HoverCardTrigger } from "@/components/ui/hover-card";

function App() {
  const [data, setData] = useState('');
  const [longUrl, setLongUrl] = useState("");

  const handleTidyUp: MouseEventHandler<HTMLButtonElement> = async () => {
    const tidyUrlResponse = await tidyUrlServiceClient.makeTidyUrl({
      longUrl
    })

    setData(tidyUrlResponse.tidyUrl)
  };

  console.log("data", data);

  return (
    <section className="bg-[#ccebe3]">
      <GithubCorner />
      <div className="h-screen flex items-center justify-center">
        <div className="w-2xs md:w-90 flex gap-4">
          <Input onChange={(e) => setLongUrl(e.target.value)} className="border-[#a5c3c2] bg-white selection:bg-white focus-visible:ring-[#a5c3c2]" placeholder="Paste Your Long Url Here" />
          <Button onClick={handleTidyUp} className="bg-white text-[#a5c3c2]">Tidy</Button>
          <HoverCard>
            <HoverCardTrigger asChild>
              <Button
                variant="link"
              >
                Grab Here!
              </Button>
            </HoverCardTrigger>
            <HoverCardContent>
              <p className="text-sm text-muted-foreground">
                {data}
              </p>
            </HoverCardContent>
          </HoverCard>
        </div>
      </div>
    </section>
  );
}

export default App;
