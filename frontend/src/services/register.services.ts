"use client";

export async function RegisterService(userData: { PhoneNumber: string; Email: string, Password: string }) {
  const res = await fetch("/user/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(userData),
  });
  return await res.json();
}
