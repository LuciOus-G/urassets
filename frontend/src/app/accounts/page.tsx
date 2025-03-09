"use client"

import {useState} from "react";
import MultiSelect from "@/components/multiSelect";
import {cn} from "@/lib/utils";

export default function Page() {
  const [formData, setFormData] = useState({
    fullName: '',
    email: '',
    password: '',
    confirmPassword: '',
    currentPassword: '',
    payDayDate: '',
    expensesCategory: '',
    incomeCategory: '',
  });

  const handleOnChangeDefault = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({...formData, [e.target.name]: e.target.value})
  }

  return (
    <div className="p-5 rounded-xl bg-white">
      <h1 className="text-xl font-bold text-black">Profile and Settings</h1>

      <div className="mt-5">
        <div className="grid grid-cols-7 gap-4 my-3">
          <label htmlFor="fullName" className="text-black">Full Name: </label>
          <div
            className="col-span-6 flex items-center justify-start w-1/2 p-1 bg-white border border-gray-200 rounded-full shadow-sm"
          >
            <input id={"fullName"} name={"fullName"} className="rounded-full w-full p-2 text-black" value={formData.fullName} onChange={handleOnChangeDefault} />
          </div>
        </div>

        <div className="grid grid-cols-7 gap-4 my-3">
          <label htmlFor="email" className="text-black">Email: </label>
          <div className="col-span-6 flex items-center justify-start w-1/2 p-1 bg-white border border-gray-200 rounded-full shadow-sm">
            <input name={"email"} onChange={handleOnChangeDefault} type={"email"} id={"email"} className="rounded-full w-full p-2 text-black" value={formData.email}/>
          </div>
        </div>

        <div className="grid grid-cols-7 gap-4 my-3">
          <label htmlFor="password" className="text-black">New Password: </label>
          <div className="col-span-6 flex items-center justify-start w-1/2 p-1 bg-white border border-gray-200 rounded-full shadow-sm">
            <input name={"password"} onChange={handleOnChangeDefault} type={"password"} id={"password"} className="rounded-full w-full p-2 text-black" value={formData.password}/>
          </div>
        </div>

        <div className="grid grid-cols-7 gap-4 my-3">
          <label htmlFor="confirmPassword" className="text-black">Confirm Password: </label>
          <div className="col-span-6 flex items-center justify-start w-1/2 p-1 bg-white border border-gray-200 rounded-full shadow-sm">
            <input name={"confirmPassword"} onChange={handleOnChangeDefault} type={"password"} id={"confirmPassword"} className="rounded-full w-full p-2 text-black" value={formData.confirmPassword}/>
          </div>
        </div>

        <div className="grid grid-cols-7 gap-4 my-3">
          <label htmlFor="currentPassword" className="text-black">Current Password: </label>
          <div className="col-span-6 flex items-center justify-start w-1/2 p-1 bg-white border border-gray-200 rounded-full shadow-sm">
            <input name={"currentPassword"} onChange={handleOnChangeDefault} type={"password"} id={"currentPassword"} className="rounded-full w-full p-2 text-black" value={formData.currentPassword}/>
          </div>
        </div>

        <div className="grid grid-cols-7 gap-4 my-3">
          <label htmlFor="payDayDate" className="text-black">Payday Date: </label>
          <div className="col-span-6 flex items-center justify-start w-1/2 p-1 bg-white border border-gray-200 rounded-full shadow-sm">
            <input name={"payDayDate"} onChange={handleOnChangeDefault} type={"date"} id={"payDayDate"} className="rounded-full w-full p-2 text-black" value={formData.payDayDate}/>
          </div>
        </div>

        <div className="grid grid-cols-7 gap-4 my-3">
          <label htmlFor="expensesCategory" className="text-black">Expenses Category: </label>
          <div className="col-span-6 flex items-center justify-start w-1/2 p-1 bg-white border border-gray-200 rounded-full shadow-sm">
            <MultiSelect BodyClass={"bg-white border border-gray-200 text-black rounded-full"}/>
          </div>
        </div>

        <div className="grid grid-cols-7 gap-4 my-3">
          <label htmlFor="incomeCategory" className="text-black">Income Category: </label>
          <div className="col-span-6 flex items-center justify-start w-1/2 p-1 bg-white border border-gray-200 rounded-full shadow-sm">
            <MultiSelect BodyClass={cn("bg-white border border-gray-200 text-black rounded-full")}/>
          </div>
        </div>
      </div>
    </div>
  )
}