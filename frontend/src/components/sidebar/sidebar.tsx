"use client"

import { CreditCard, Coins, LayoutGrid, RefreshCcw, Goal, Settings, User, Wallet } from "lucide-react"
import { useState } from "react"
import { SidebarItem } from "@/components/sidebar/sidebar-item"
import {useRouter} from "next/navigation";

interface SidebarProps {
  onNavigate?: (route: string) => void
}

export function Sidebar({ onNavigate }: SidebarProps) {
  const [activeRoute, setActiveRoute] = useState("overview")
  const router = useRouter()

  const handleNavigation = (route: string) => {
    setActiveRoute(route)
    if (onNavigate) {
      onNavigate(route)
    }
  }

  return (
    <div className="w-64 bg-white p-6 flex flex-col relative fixed h-full top-0 left-0 shadow-lg">
      <div className="flex items-center mb-10 justify-center align-center">
        <h1 className="text-xl font-bold text-black">UrAssets</h1>
      </div>

      <nav className="flex-1 space-y-1">
        <SidebarItem
          icon={LayoutGrid}
          label="Summary"
          isActive={activeRoute === "summary"}
          onClick={() => {
              handleNavigation("summary")
              router.push("/summary")
            }
          }
        />

        <SidebarItem
          icon={User}
          label="Profile"
          isActive={activeRoute === "accounts"}
          onClick={() => {
              handleNavigation("accounts")
              router.push("/accounts")
            }
          }
        />

        <SidebarItem
          icon={CreditCard}
          label="Bank Accounts"
          isActive={activeRoute === "bank"}
          onClick={() => {
            handleNavigation("bank")
            router.push("/bank-accounts")
          }}
        />

        <SidebarItem
          icon={Coins}
          label="Budgeting"
          isActive={activeRoute === "budget"}
          onClick={() => {
            handleNavigation("budget")
            router.push("/budgeting")
          }}
        />

        <SidebarItem
          icon={Goal}
          label="Goals"
          isActive={activeRoute === "goals"}
          onClick={() => {
            handleNavigation("goals")
            router.push("/goals")
          }}
        />

        <SidebarItem
          icon={RefreshCcw}
          label="Transactions"
          isActive={activeRoute === "transactions"}
          onClick={() => {
            handleNavigation("transactions")
            router.push("/transactions")
          }}
        />

        <SidebarItem
          icon={Wallet}
          label="Assets"
          isActive={activeRoute === "assets"}
          onClick={() => handleNavigation("assets")}
        />
      </nav>

      <SidebarItem
        icon={Settings}
        label="Settings"
        isActive={activeRoute === "settings"}
        onClick={() => handleNavigation("settings")}
      />
    </div>
  )
}

