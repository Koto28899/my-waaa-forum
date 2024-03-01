import Link from "next/link";

function NotFoundPage() {
  return (
    <div className="not-found">
      <h1>404 Not Found</h1>
      <p>Could not find requested resource</p>
      <Link href="/">Return Home</Link>
    </div>
  )
}

export default NotFoundPage;