"use client"

import Image from "next/image"
import {Ban, Check, KeyRound, Mail, Phone} from "lucide-react"
import React, {useEffect, useState} from "react";
import {LoginService} from "@/services/login.services";
import {useRouter} from "next/navigation";
import toast from "react-hot-toast";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import {cn} from "@/lib/utils";
import Error from "next/error";


const RegisterSchema = z.object({
    email: z.string().email("Invalid email format"),
    phoneNumber: z.string()
      .min(8, "min length is 8")
      .max(12, "max length is 12")
      .regex(/^\d+$/, "Phone number must contain only digits"),
    password: z.string(),
});
type RegisterFormData = z.infer<typeof RegisterSchema>;

export default function Login() {
    const router = useRouter()

    useEffect(() => {
        if (sessionStorage.getItem("UserData")){
            router.push("/dashboard/summary")
        }
    }, []);

    // initiate form data
    const { register, handleSubmit, formState: { errors, isSubmitting }} = useForm<RegisterFormData>({
        resolver: zodResolver(RegisterSchema),
        defaultValues: {
            email: "",
            password: "",
        },
    });

    const handleOnLogin = (data: RegisterFormData) => {
        toast.promise(
            LoginService({email: data.email, password: data.password}).then((r) => {
                if (r.status === 200) {
                    sessionStorage.setItem("UserData", JSON.stringify(r))
                    router.push("/dashboard/summary")
                } else {
                    throw new Error(r.error);
                }
            }).catch((err) => {
                throw new Error(err.props);
            }),
            {
                loading: 'Logging...',
                success: <b>Login</b>,
                error: (err) => <b>{err.props}</b>,
            }
        );
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
                    <h1 className="text-3xl font-bold mb-2">Welcome</h1>
                    <p className="text-gray-500 mb-8">Please enter Your details</p>

                    <div className="flex mb-8 border-b">
                        <button className="pb-2 px-4 font-medium border-b-2 border-black">Register</button>
                    </div>

                    <form onSubmit={handleSubmit(handleOnLogin)}>
                        <div className="relative mb-6">
                            <div className="absolute inset-y-0 left-0 flex items-center pl-4 pointer-events-none">
                                <Mail className="h-5 w-5 text-gray-400" />
                            </div>
                            <input
                                type="email"
                                className="w-full pl-12 pr-10 py-3 border rounded-lg"
                                placeholder="Email Address"
                                {...register("email")}
                            />
                            <div className="absolute inset-y-0 right-0 flex items-center pr-3">
                                <div className="bg-green-100 p-1 rounded-full">
                                    <Check className="h-4 w-4 text-green-500" />
                                </div>
                            </div>
                        </div>

                        <div className="relative mb-6">
                            <div className="absolute inset-y-0 left-0 flex items-center pl-4 pointer-events-none">
                                <Phone className="h-5 w-5 text-gray-400" />
                            </div>
                            <input
                              type="text"
                              className="w-full pl-12 pr-10 py-3 border rounded-lg"
                              placeholder="Phone Number"
                              {...register("phoneNumber")}
                            />
                            <div className="absolute inset-y-0 right-0 flex items-center pr-3">
                                <div className="bg-green-100 p-1 rounded-full">
                                    {errors.phoneNumber ? <Ban className="h-4 w-4 text-green-500"/> : <Check className="h-4 w-4 text-green-500"/>}
                                </div>
                            </div>
                        </div>

                        <div className="relative mb-6">
                            <div className="absolute inset-y-0 left-0 flex items-center pl-4 pointer-events-none">
                                <KeyRound className="h-5 w-5 text-gray-400" />
                            </div>
                            <input
                                type="password"
                                className="w-full pl-12 pr-10 py-3 border rounded-lg"
                                placeholder="Password"
                                {...register("password")}
                            />
                            <div className="absolute inset-y-0 right-0 flex items-center pr-3">
                                <div className="bg-green-100 p-1 rounded-full">
                                    <Check className="h-4 w-4 text-green-500" />
                                </div>
                            </div>
                        </div>

                        <button type="submit" className="hover:cursor-pointer bg-blue-600 text-white py-3 w-full rounded-lg font-medium mb-8" disabled={isSubmitting}>{isSubmitting ? "Registering..." : "Register"}</button>
                    </form>
                    <div className="flex items-center gap-4 mb-8">
                        <div className="h-px bg-gray-200 flex-1"></div>
                        <span className="text-gray-400 text-sm">Have an account?</span>
                        <div className="h-px bg-gray-200 flex-1"></div>
                    </div>

                    <div className="flex justify-center gap-4 mb-12">
                        <button onClick={() => {router.push("/login")}} type="button" className="hover:cursor-pointer bg-zinc-600 text-white py-3 w-full rounded-lg font-medium mb-8">Login Here</button>
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

