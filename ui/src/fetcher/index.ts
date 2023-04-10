export default async function fetcher(
  url: string,
  { arg }: { arg: { long_url: string } }
) {
  return fetch(url, {
    method: "POST",
    body: JSON.stringify(arg),
  }).then((res) => res.json());
}
