"use client"

import {useState} from "react";
import MultiSelect from "@/components/multiSelect";
import {cn} from "@/lib/utils";
import BookingTable from "@/components/table/bookingTable";
import Image from "next/image";
import {MoreHorizontal} from "lucide-react";

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

      <div className={"my-16"}>
        <h1 className={"text-bold text-2xl mb-5"}>Budgeting</h1>
        <div className={"flex justify-between items-center"}>
          <h1>Add Income</h1>
          <button  type={"button"} className={"p-2 bg-green-400 rounded-lg hover:cursor-pointer text-green-900"}>+ Add Item</button>
        </div>

        <div className={cn("overflow-x-auto")}>
          <table className="min-w-full divide-y divide-gray-200">
            <thead>
            <tr className="text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              <th className="px-4 py-3 w-10">
                <input type="checkbox" className="h-4 w-4 rounded border-gray-300 text-primary focus:ring-primary" />
              </th>
              <th className="px-4 py-3">Booking No</th>
              <th className="px-4 py-3">Name of Guest</th>
              <th className="px-4 py-3">Source</th>
              <th className="px-4 py-3">Date</th>
              <th className="px-4 py-3">Email</th>
              <th className="px-4 py-3">Mobile Number</th>
              <th className="px-4 py-3">Amount</th>
              <th className="px-4 py-3">Status</th>
              <th className="px-4 py-3 w-10"></th>
            </tr>
            </thead>
            <tbody className="bg-white divide-y divide-gray-200">
              <tr key="1" className="hover:bg-gray-50">
                <td className="px-4 py-3 whitespace-nowrap">
                  <input type="checkbox" className="h-4 w-4 rounded border-gray-300 text-primary focus:ring-primary" />
                </td>
                <td className="px-4 py-3 whitespace-nowrap text-sm font-medium text-gray-900">1</td>
                <td className="px-4 py-3 whitespace-nowrap">
                  <div className="flex items-center">
                    <div className="ml-3">
                      <div className="text-sm font-medium text-gray-900">Income</div>
                    </div>
                  </div>
                </td>
                <td className="px-4 py-3 whitespace-nowrap text-sm text-gray-500">BCA</td>
                <td className="px-4 py-3 whitespace-nowrap text-sm text-gray-500">123</td>
                <td className="px-4 py-3 whitespace-nowrap text-sm text-gray-500">123</td>
                <td className="px-4 py-3 whitespace-nowrap text-sm text-gray-500">123</td>
                <td className="px-4 py-3 whitespace-nowrap text-sm font-medium text-gray-900">123</td>
                <td className="px-4 py-3 whitespace-nowrap">
                <span
                  className={`inline-flex rounded-full px-3 py-1 text-xs font-medium ${
                    "booking.status" === "Cash - Paid"
                      ? "bg-green-100 text-green-800"
                      : "booking.status" === "Online - Paid"
                        ? "bg-green-100 text-green-800"
                        : "bg-yellow-100 text-yellow-800"
                  }`}
                >
                  "booking.status"
                </span>
                </td>
                <td className="px-4 py-3 whitespace-nowrap text-right text-sm font-medium">
                  <button className="text-gray-400 hover:text-gray-500">
                    <MoreHorizontal className="h-5 w-5" />
                  </button>
                </td>
              </tr>
            </tbody>
          </table>

          <div className="flex items-center justify-between border-t border-gray-200 px-4 py-3 sm:px-6">
            <div className="flex flex-1 justify-between sm:hidden">
              <button className="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">
                Previous
              </button>
              <button className="relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">
                Next
              </button>
            </div>
            <div className="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
              <div></div>
              <div>
                <nav className="isolate inline-flex -space-x-px rounded-md shadow-sm" aria-label="Pagination">
                  <button className="relative inline-flex items-center rounded-l-md px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20">
                    <span className="sr-only">Previous</span>
                    <svg
                      className="h-5 w-5"
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                      aria-hidden="true"
                    >
                      <path
                        fillRule="evenodd"
                        d="M12.79 5.23a.75.75 0 01-.02 1.06L8.832 10l3.938 3.71a.75.75 0 11-1.04 1.08l-4.5-4.25a.75.75 0 010-1.08l4.5-4.25a.75.75 0 011.06.02z"
                        clipRule="evenodd"
                      />
                    </svg>
                    Previous
                  </button>
                  <button
                    aria-current="page"
                    className="relative z-10 inline-flex items-center bg-green-600 px-4 py-2 text-sm font-medium text-white focus:z-20"
                  >
                    1
                  </button>
                  <button className="relative inline-flex items-center px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20">
                    2
                  </button>
                  <span className="relative inline-flex items-center px-4 py-2 text-sm font-medium text-gray-700">...</span>
                  <button className="relative inline-flex items-center px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20">
                    10
                  </button>
                  <button className="relative inline-flex items-center px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20">
                    12
                  </button>
                  <button className="relative inline-flex items-center px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20">
                    13
                  </button>
                  <button className="relative inline-flex items-center px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20">
                    14
                  </button>
                  <button className="relative inline-flex items-center rounded-r-md px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 focus:z-20">
                    Next
                    <svg
                      className="h-5 w-5"
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                      aria-hidden="true"
                    >
                      <path
                        fillRule="evenodd"
                        d="M7.21 14.77a.75.75 0 01.02-1.06L11.168 10 7.23 6.29a.75.75 0 111.04-1.08l4.5 4.25a.75.75 0 010 1.08l-4.5 4.25a.75.75 0 01-1.06-.02z"
                        clipRule="evenodd"
                      />
                    </svg>
                  </button>
                </nav>
              </div>
            </div>
          </div>
        </div>

      </div>

    </div>
  )
}