"use client";

export async function LoginService(userData: { email: string; password: string }) {
  const res = await fetch("/user/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(userData), // Convert userData to JSON string
  });

  return await res.json();
}
