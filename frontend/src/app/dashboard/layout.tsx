"use client";

import React, {useEffect, useState} from "react";
import { useRouter } from "next/navigation";
import {Sidebar} from "@/components/sidebar/sidebar";
import {PuffLoader} from "react-spinners";

export default function DashboardLayout({ children }: { children: React.ReactNode }) {
    const router = useRouter();
    const [isLoading, setIsLoading] = useState(false)

    useEffect(() => {
        const token = sessionStorage.getItem("token");

        if (!token) {
        } else {
            setIsLoading(false)
        }
    }, [router]);

    if (isLoading) {
        return (
            <div className="flex h-screen items-center justify-center">
                <PuffLoader color="#36d7b7" size={80} />
            </div>
        );
    }

    return <>
        <div className="flex h-screen bg-neutral-100">
            <Sidebar/>
            <div className="p-5 h-full w-full overflow-y-auto">
                {children}
            </div>
        </div>
    </>;
}
