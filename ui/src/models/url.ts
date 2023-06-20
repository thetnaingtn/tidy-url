import { validateURL } from "../util";

class LongURL {
  private longUrl: string;

  public constructor(url: string) {
    this.longUrl = url;
  }

  get isURLValid(): boolean {
    return validateURL(this.longUrl);
  }

  get error(): string {
    if (!this.isURLValid) {
      return "The input URL is invalid.";
    }
    return "";
  }
}

export { LongURL };
