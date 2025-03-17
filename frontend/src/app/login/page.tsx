"use client"

import Image from "next/image"
import { Check, Mail } from "lucide-react"

export default function Login() {
    function onLoginSubmit() {

    }

    return (
        <main className="flex min-h-screen">
            {/* Left Section */}
            <div className="flex flex-col w-full lg:w-1/2 p-8 md:p-16 lg:p-24">
                <div className="mb-16">
                    <div className="flex items-center gap-2">
                        <div className="w-6 h-6">
                            <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path d="M12 2L2 7L12 12L22 7L12 2Z" fill="black" />
                                <path
                                    d="M2 17L12 22L22 17"
                                    stroke="black"
                                    strokeWidth="2"
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                />
                                <path
                                    d="M2 12L12 17L22 12"
                                    stroke="black"
                                    strokeWidth="2"
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                />
                            </svg>
                        </div>
                        <span className="font-semibold text-xl">SmartSave</span>
                    </div>
                </div>

                <div className="flex-1 flex flex-col justify-center max-w-md mx-auto w-full">
                    <h1 className="text-3xl font-bold mb-2">Welcome Back</h1>
                    <p className="text-gray-500 mb-8">Welcome Back, Please enter Your details</p>

                    <div className="flex mb-8 border-b">
                        <button className="pb-2 px-4 font-medium border-b-2 border-black">Sign In</button>
                        <button className="pb-2 px-4 text-gray-500">Signup</button>
                    </div>

                    <div className="relative mb-6">
                        <div className="absolute inset-y-0 left-0 flex items-center pl-4 pointer-events-none">
                            <Mail className="h-5 w-5 text-gray-400" />
                        </div>
                        <input
                            type="email"
                            className="w-full pl-12 pr-10 py-3 border rounded-lg"
                            placeholder="Email Address"
                            defaultValue="ialirezamp@gmail.com"
                        />
                        <div className="absolute inset-y-0 right-0 flex items-center pr-3">
                            <div className="bg-green-100 p-1 rounded-full">
                                <Check className="h-4 w-4 text-green-500" />
                            </div>
                        </div>
                    </div>

                    <button onClick={onLoginSubmit} className="hover:cursor-pointer bg-blue-600 text-white py-3 rounded-lg font-medium mb-8">Continue</button>

                    <div className="flex items-center gap-4 mb-8">
                        <div className="h-px bg-gray-200 flex-1"></div>
                        <span className="text-gray-400 text-sm">Or Continue With</span>
                        <div className="h-px bg-gray-200 flex-1"></div>
                    </div>

                    <div className="flex justify-center gap-4 mb-12">
                        <button className="w-12 h-12 flex items-center justify-center rounded-full border">
                            <svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 0 24 24" width="24">
                                <path
                                    d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
                                    fill="#4285F4"
                                />
                                <path
                                    d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
                                    fill="#34A853"
                                />
                                <path
                                    d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
                                    fill="#FBBC05"
                                />
                                <path
                                    d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
                                    fill="#EA4335"
                                />
                                <path d="M1 1h22v22H1z" fill="none" />
                            </svg>
                        </button>
                        <button className="w-12 h-12 flex items-center justify-center rounded-full border bg-black">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="white">
                                <path d="M17.05 20.28c-.98.95-2.05.88-3.08.45-1.09-.46-2.09-.48-3.24 0-1.44.62-2.2.44-3.06-.45C2.08 14.34 3.23 6.04 9.62 5.7c1.5.06 2.52.95 3.33 1 1.32-.17 2.57-1.02 3.9-.87 1.65.2 2.88.96 3.64 2.4-3.3 2.14-2.77 6.16.35 7.72-.64 1.93-1.47 3.85-3.79 4.33z" />
                                <path d="M12.77 5.53c-.15-2.37 1.74-4.5 3.94-4.53.26 2.05-1.85 4.67-3.94 4.53z" />
                            </svg>
                        </button>
                        <button className="w-12 h-12 flex items-center justify-center rounded-full border bg-blue-600">
                            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="white">
                                <path d="M12 2.04c-5.5 0-10 4.49-10 10.02 0 5 3.66 9.15 8.44 9.9v-7H7.9v-2.9h2.54V9.85c0-2.51 1.49-3.89 3.78-3.89 1.09 0 2.23.19 2.23.19v2.47h-1.26c-1.24 0-1.63.77-1.63 1.56v1.88h2.78l-.45 2.9h-2.33v7a10 10 0 0 0 8.44-9.9c0-5.53-4.5-10.02-10-10.02z" />
                            </svg>
                        </button>
                    </div>

                    <p className="text-center text-gray-500 text-sm leading-relaxed">
                        Join the millions of smart investors who trust us to manage their finances. Log in to access your
                        personalized dashboard, track your portfolio performance, and make informed investment decisions.
                    </p>
                </div>
            </div>

            {/* Right Section - Image */}
            <div className="hidden lg:block w-1/2 bg-blue-50 relative">
                <div className="absolute inset-0 flex items-center justify-center p-12">
                    <Image
                        src="/placeholder.svg?height=500&width=500"
                        alt="Safe illustration"
                        width={500}
                        height={500}
                        className="object-contain"
                        priority
                    />
                </div>
            </div>
        </main>
    )
}

