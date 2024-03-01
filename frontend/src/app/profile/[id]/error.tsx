'use client'
 
import Link from 'next/link'
import { useEffect } from 'react'
 
export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string }
  reset: () => void
}) {
  useEffect(() => {
    console.error(error)
  }, [error])

  return (
    <div className='error'>
      <h1>Something went wrong!</h1>
      <p>Could not get requested resource</p>
      <p>Error Detail: <span>{error.message}</span></p>
      <Link className='button' href="/">Return Home</Link>
      <button className='button' onClick={() => reset()}>
        Try again
      </button>
    </div>
  )
}