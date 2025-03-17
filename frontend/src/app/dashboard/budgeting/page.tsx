"use client"

import { useState } from "react"
import {Bell, ChevronDown, Eye, Plus} from "lucide-react";
import {router} from "next/client";
import {useRouter} from "next/navigation";
import ColumnTable from "@/components/column";

export default function App() {
    const router = useRouter()

    const [accountsExpanded, setAccountsExpanded] = useState(true)
    const [spendingExpanded, setSpendingExpanded] = useState(true)

    return (
        <ColumnTable/>
    )
}

