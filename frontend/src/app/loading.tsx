"use client";

import { PuffLoader } from "react-spinners";

export default function Loading() {
    return (
        <div className="flex h-screen items-center justify-center">
            <PuffLoader color="#36d7b7" size={80} />
        </div>
    );
}
