"use client"

import { useState } from "react"
import {Bell, ChevronDown, Eye, Plus} from "lucide-react";
import {router} from "next/client";
import {useRouter} from "next/navigation";

export default function App() {
  const router = useRouter()

  const [accountsExpanded, setAccountsExpanded] = useState(true)
  const [spendingExpanded, setSpendingExpanded] = useState(true)

  return (
    <div className="overflow-auto">
      {/* Header */}
      <div className="bg-green-400 rounded-3xl p-8 mb-4">
        <div className="flex justify-between mb-6">
          <h1 className="text-4xl font-bold text-slate-950">
            Welcome Back <br/> Wilson!
          </h1>
          <div>
            <p className={"text-black text-2xl"}>Money Status:</p>
            <span className={"text-green-900 text-center"}>Excellent!</span>
          </div>
        </div>

        <div className="flex justify-between items-end text-black">
          <div>
            <p className="text-gray-700 mb-1">Total Wealth</p>
            <h2 className="text-4xl font-bold">$547,197.72</h2>
          </div>
          <div>
            <p className="text-gray-700 mb-1">Total Loan</p>
            <h2 className="text-4xl font-bold">$23,456.00</h2>
          </div>
        </div>
      </div>

      {/* Accounts and Cards Section */}
      <div className="mb-4">
        <button
          className="w-full bg-orange-200 hover:bg-orange-300 p-4 rounded-t-2xl flex justify-between items-center"
          onClick={() => setAccountsExpanded(!accountsExpanded)}
        >
          <h2 className="text-xl font-medium text-orange-800">Accounts and cards</h2>
          <ChevronDown
            className={`h-6 w-6 text-orange-800 transition-transform ${accountsExpanded ? "rotate-180" : ""}`}
          />
        </button>
        {accountsExpanded && (
          <div className="p-4 bg-white rounded-b-2xl border border-orange-200 border-t-0">
            {/* Account content would go here */}
          </div>
        )}
      </div>

      {/* Spending Section */}
      <div className="mb-4">
        <button
          className="w-full bg-blue-200 hover:bg-blue-300 p-4 rounded-t-2xl flex justify-between items-center"
          onClick={() => setSpendingExpanded(!spendingExpanded)}
        >
          <h2 className="text-xl font-medium text-blue-800">Your spending</h2>
          <ChevronDown
            className={`h-6 w-6 text-blue-800 transition-transform ${spendingExpanded ? "rotate-180" : ""}`}
          />
        </button>
        {spendingExpanded && (
          <div className="p-4 bg-white rounded-b-2xl border border-blue-200 border-t-0">
            {/* Spending content would go here */}
          </div>
        )}
      </div>

      {/* Transactions Section */}
      <div>
        <div className="flex justify-between align-center items-center items-baseline mb-4">
          <h2 className="text-2xl font-bold text-black">Recent Transactions</h2>
          <span onClick={() => router.push("/transactions")} className="ml-2 text-sm text-blue-500 hover:cursor-pointer">View All</span>
        </div>

        <div className="space-y-4">
          {/* Apple Transaction */}
          <div className="flex items-center justify-between bg-white p-4 rounded-xl">
            <div className="flex items-center">
              <div className="w-10 h-10 bg-gray-100 rounded-lg flex items-center justify-center mr-4">
                <svg viewBox="0 0 24 24" width="24" height="24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path
                    d="M12.152 6.896c-.948 0-2.415-1.078-3.96-1.04-2.04.027-3.91 1.183-4.961 3.014-2.117 3.675-.54 9.103 1.519 12.09 1 1.468 2.2 3.105 3.773 3.043 1.52-.065 2.09-.973 3.926-.973 1.836 0 2.35.973 3.958.946 1.635-.026 2.668-1.49 3.667-2.96 1.156-1.687 1.635-3.325 1.662-3.415-.039-.013-3.182-1.221-3.22-4.857-.026-3.04 2.48-4.494 2.597-4.559-1.429-2.09-3.623-2.324-4.39-2.376-2-.156-3.675 1.09-4.571 1.09z"
                    fill="#000"
                  />
                  <path
                    d="M14.128 3.973c.838-1.016 1.39-2.428 1.235-3.83-1.196.052-2.65.8-3.514 1.816-.772.89-1.443 2.324-1.26 3.7 1.326.103 2.7-.677 3.539-1.686z"
                    fill="#000"
                  />
                </svg>
              </div>
              <div>
                <p className="font-medium">Apple store</p>
                <p className="text-sm text-gray-500">10 Sep 2023 at 3:12 PM</p>
              </div>
            </div>
            <span className="font-bold text-red-600">-$2,456.67</span>
          </div>

          {/* Stripe Transaction */}
          <div className="flex items-center justify-between bg-white p-4 rounded-xl">
            <div className="flex items-center">
              <div className="w-10 h-10 bg-gray-100 rounded-lg flex items-center justify-center mr-4">
                <svg viewBox="0 0 24 24" width="24" height="24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path
                    d="M13.479 9.883c0-.773-.626-1.076-1.636-1.076-.626 0-1.328.152-1.328.152l-.177.76s.677-.152 1.303-.152c.455 0 .657.126.657.379 0 .202-.05.354-.05.354h-.48c-.753 0-1.53.303-1.53 1.353 0 .808.53 1.01 1.05 1.01.783 0 1.177-.505 1.177-.505l-.025.429h.884s.126-.783.126-1.01l.025-1.694zm-.884 1.01s-.126.556-.733.556c-.303 0-.404-.253-.404-.379 0-.278.152-.606.884-.606h.253v.43z"
                    fill="#635BFF"
                  />
                  <path
                    d="M15.317 8.845h-1.01l.025-.177s.354-.76 1.01-.76c.657 0 .834.556.834.834 0 .278-.05.556-.05.556h.657c.556 0 .505.556.505.556l-.05.303h-.657l-.227 1.313c-.025.177.05.253.227.253.126 0 .253-.05.253-.05l-.05.556s-.354.101-.657.101c-.303 0-.657-.177-.631-.657l.227-1.515h-.404l.101-.556h.404l.05-.303c0-.025 0-.05-.025-.05h-.126zm-3.788-1.01h-.96l-.05.303h.96l.05-.303zm-.177 1.01h-.96l-.505 2.98h.96l.505-2.98zm-1.49 0h-.96l-.354 2.146s-.126.657-.48.657c-.227 0-.278-.177-.278-.177l-.152.884s.177.126.48.126c.303 0 .657-.177.884-.657l.657-2.98h-.808zm-2.45 1.667s.404-.404.404-.884c0-.48-.404-.783-.884-.783-.48 0-.884.303-.884.783 0 .48.404.884.884.884.48 0 .48 0 .48 0zm.177.177h-.354c-.227 0-.404-.05-.556-.126l-.278 1.667h-.96l.707-4.192h.96l-.126.733s.354-.404.834-.404c.48 0 .884.404.884 1.01 0 1.212-1.111 1.313-1.111 1.313z"
                    fill="#635BFF"
                  />
                </svg>
              </div>
              <div>
                <p className="font-medium">Stripe payment</p>
                <p className="text-sm text-gray-500">09 Sep 2023 at 2:10 PM</p>
              </div>
            </div>
            <span className="font-bold text-green-600">+$346.00</span>
          </div>

          {/* Meta Transaction */}
          <div className="flex items-center justify-between bg-white p-4 rounded-xl">
            <div className="flex items-center">
              <div className="w-10 h-10 bg-gray-100 rounded-lg flex items-center justify-center mr-4">
                <svg viewBox="0 0 24 24" width="24" height="24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path
                    d="M12 2C6.477 2 2 6.477 2 12c0 4.991 3.657 9.128 8.438 9.879V14.89h-2.54V12h2.54V9.797c0-2.506 1.492-3.89 3.777-3.89 1.094 0 2.238.195 2.238.195v2.46h-1.26c-1.243 0-1.63.771-1.63 1.562V12h2.773l-.443 2.89h-2.33v6.989C18.343 21.129 22 16.99 22 12c0-5.523-4.477-10-10-10z"
                    fill="#1877F2"
                  />
                </svg>
              </div>
              <div>
                <p className="font-medium">Meta subscription</p>
                <p className="text-sm text-gray-500">07 Sep 2023 at 10:02 AM</p>
              </div>
            </div>
            <span className="font-bold text-red-600">-$179.88</span>
          </div>
          <div className="flex items-center justify-between bg-white p-4 rounded-xl">
            <div className="flex items-center">
              <div className="w-10 h-10 bg-gray-100 rounded-lg flex items-center justify-center mr-4">
                <svg viewBox="0 0 24 24" width="24" height="24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path
                    d="M12 2C6.477 2 2 6.477 2 12c0 4.991 3.657 9.128 8.438 9.879V14.89h-2.54V12h2.54V9.797c0-2.506 1.492-3.89 3.777-3.89 1.094 0 2.238.195 2.238.195v2.46h-1.26c-1.243 0-1.63.771-1.63 1.562V12h2.773l-.443 2.89h-2.33v6.989C18.343 21.129 22 16.99 22 12c0-5.523-4.477-10-10-10z"
                    fill="#1877F2"
                  />
                </svg>
              </div>
              <div>
                <p className="font-medium">Meta subscription</p>
                <p className="text-sm text-gray-500">07 Sep 2023 at 10:02 AM</p>
              </div>
            </div>
            <span className="font-bold text-red-600">-$179.88</span>
          </div>
          <div className="flex items-center justify-between bg-white p-4 rounded-xl">
            <div className="flex items-center">
              <div className="w-10 h-10 bg-gray-100 rounded-lg flex items-center justify-center mr-4">
                <svg viewBox="0 0 24 24" width="24" height="24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path
                    d="M12 2C6.477 2 2 6.477 2 12c0 4.991 3.657 9.128 8.438 9.879V14.89h-2.54V12h2.54V9.797c0-2.506 1.492-3.89 3.777-3.89 1.094 0 2.238.195 2.238.195v2.46h-1.26c-1.243 0-1.63.771-1.63 1.562V12h2.773l-.443 2.89h-2.33v6.989C18.343 21.129 22 16.99 22 12c0-5.523-4.477-10-10-10z"
                    fill="#1877F2"
                  />
                </svg>
              </div>
              <div>
                <p className="font-medium">Meta subscription</p>
                <p className="text-sm text-gray-500">07 Sep 2023 at 10:02 AM</p>
              </div>
            </div>
            <span className="font-bold text-red-600">-$179.88</span>
          </div>
        </div>
      </div>
    </div>
  )
}

