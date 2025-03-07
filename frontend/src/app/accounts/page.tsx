"use client"
import {Plus} from "lucide-react";

export default function Page() {

  return (
    <div>
      <div className="flex justify items-center">
        <button className="flex justify items-center bg-purple-400 p-2 rounded-lg">
          <Plus className="text-purple-800"/>
          <h1 className="text-sm font-bold text-purple-800">Add bank accounts</h1>
        </button>
      </div>
    </div>
  )
}

