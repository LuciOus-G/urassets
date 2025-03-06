"use client"

import type { LucideIcon } from "lucide-react"

interface SidebarItemProps {
  icon: LucideIcon
  label: string
  isActive?: boolean
  onClick?: () => void
}

export function SidebarItem({ icon: Icon, label, isActive = false, onClick }: SidebarItemProps) {
  return (
    <button
      className={`w-full flex items-center p-3 rounded-lg ${
        isActive ? "bg-gray-100 text-gray-800" : "text-gray-600 hover:bg-gray-50"
      }`}
      onClick={onClick}
    >
      <Icon className="h-5 w-5 mr-3" />
      <span>{label}</span>
    </button>
  )
}

