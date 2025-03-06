"use client"

import { CreditCard, LayoutGrid, RefreshCcw, Settings, User, Wallet } from "lucide-react"
import { useState } from "react"
import { SidebarItem } from "@/components/sidebar/sidebar-item"

interface SidebarProps {
  onNavigate?: (route: string) => void
}

export function Sidebar({ onNavigate }: SidebarProps) {
  const [activeRoute, setActiveRoute] = useState("overview")

  const handleNavigation = (route: string) => {
    setActiveRoute(route)
    if (onNavigate) {
      onNavigate(route)
    }
  }

  return (
    <div className="w-64 bg-white p-6 flex flex-col">
      <div className="flex items-center mb-10">
        <h1 className="text-xl font-bold">Menedz___her</h1>
        <span className="ml-2 text-xs text-white bg-gray-800 px-2 py-0.5 rounded">PRO</span>
      </div>

      <nav className="flex-1 space-y-1">
        <SidebarItem
          icon={LayoutGrid}
          label="Overview"
          isActive={activeRoute === "overview"}
          onClick={() => handleNavigation("overview")}
        />

        <SidebarItem
          icon={User}
          label="Accounts"
          isActive={activeRoute === "accounts"}
          onClick={() => handleNavigation("accounts")}
        />

        <SidebarItem
          icon={CreditCard}
          label="Cards"
          isActive={activeRoute === "cards"}
          onClick={() => handleNavigation("cards")}
        />

        <SidebarItem
          icon={RefreshCcw}
          label="Transactions"
          isActive={activeRoute === "transactions"}
          onClick={() => handleNavigation("transactions")}
        />

        <SidebarItem
          icon={Wallet}
          label="Exchange"
          isActive={activeRoute === "exchange"}
          onClick={() => handleNavigation("exchange")}
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

