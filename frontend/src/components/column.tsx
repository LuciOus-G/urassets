import { MoreHorizontal } from "lucide-react"
import Image from "next/image"

// Define the booking data structure
interface Booking {
    id: string
    name: string
    avatar: string
    source: string
    date: string
    email: string
    mobile: string
    amount: string
    status: "Cash - Paid" | "Online - Paid" | "Pending"
}

// Sample data based on the image
const bookings: Booking[] = [
    {
        id: "#102112",
        name: "Raiders Kae",
        avatar: "/placeholder.svg?height=32&width=32",
        source: "Font Desks",
        date: "20/12/2024",
        email: "debra.h@gmail.com",
        mobile: "+88 01766703570",
        amount: "$1400.00",
        status: "Cash - Paid",
    },
    {
        id: "#102111",
        name: "Jassy Mac",
        avatar: "/placeholder.svg?height=32&width=32",
        source: "Web Reservation",
        date: "19/12/2024",
        email: "mark@gmail.com",
        mobile: "+88 01966703578",
        amount: "$400.00",
        status: "Cash - Paid",
    },
    {
        id: "#102110",
        name: "Macculam Brad",
        avatar: "/placeholder.svg?height=32&width=32",
        source: "Font Desks",
        date: "18/12/2024",
        email: "brad@gmail.com",
        mobile: "+88 01666703570",
        amount: "$9800.00",
        status: "Pending",
    },
    {
        id: "#102109",
        name: "Jhoney Cark",
        avatar: "/placeholder.svg?height=32&width=32",
        source: "Group Reservation",
        date: "17/12/2024",
        email: "cark@gmail.com",
        mobile: "+88 01966703570",
        amount: "$2400.00",
        status: "Online - Paid",
    },
    {
        id: "#102108",
        name: "Jhon Brak",
        avatar: "/placeholder.svg?height=32&width=32",
        source: "Font Desks",
        date: "16/12/2024",
        email: "jhon@gmail.com",
        mobile: "+88 01766703544",
        amount: "$1400.00",
        status: "Pending",
    },
    {
        id: "#102107",
        name: "Red Alone",
        avatar: "/placeholder.svg?height=32&width=32",
        source: "Group Reservation",
        date: "15/12/2024",
        email: "redh@gmail.com",
        mobile: "+88 01766703533",
        amount: "$980.00",
        status: "Cash - Paid",
    },
    {
        id: "#102106",
        name: "Sude Rayden",
        avatar: "/placeholder.svg?height=32&width=32",
        source: "Web Reservation",
        date: "13/12/2024",
        email: "raydan@gmail.com",
        mobile: "+88 01766703570",
        amount: "$654.00",
        status: "Cash - Paid",
    },
    {
        id: "#102105",
        name: "Mark Alone",
        avatar: "/placeholder.svg?height=32&width=32",
        source: "Font Desks",
        date: "11/12/2024",
        email: "mark@gmail.com",
        mobile: "+88 01766703599",
        amount: "$187.00",
        status: "Cash - Paid",
    },
    {
        id: "#102104",
        name: "Mukkaram Bari",
        avatar: "/placeholder.svg?height=32&width=32",
        source: "Web Reservation",
        date: "10/12/2024",
        email: "fateh@gmail.com",
        mobile: "+88 01766703588",
        amount: "$1870.00",
        status: "Pending",
    },
    {
        id: "#102103",
        name: "Shahriyar Khan",
        avatar: "/placeholder.svg?height=32&width=32",
        source: "Font Desks",
        date: "08/12/2024",
        email: "khanh@gmail.com",
        mobile: "+88 01766703570",
        amount: "$1900.00",
        status: "Online - Paid",
    },
]

export default function ColumnTable() {
    return (
        <div className="overflow-x-auto">
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
                {bookings.map((booking) => (
                    <tr key={booking.id} className="hover:bg-gray-50">
                        <td className="px-4 py-3 whitespace-nowrap">
                            <input type="checkbox" className="h-4 w-4 rounded border-gray-300 text-primary focus:ring-primary" />
                        </td>
                        <td className="px-4 py-3 whitespace-nowrap text-sm font-medium text-gray-900">{booking.id}</td>
                        <td className="px-4 py-3 whitespace-nowrap">
                            <div className="flex items-center">
                                <div className="h-8 w-8 flex-shrink-0 rounded-full overflow-hidden bg-gray-100">
                                    <Image
                                        src={booking.avatar || "/placeholder.svg"}
                                        alt={booking.name}
                                        width={32}
                                        height={32}
                                        className="h-8 w-8 rounded-full"
                                    />
                                </div>
                                <div className="ml-3">
                                    <div className="text-sm font-medium text-gray-900">{booking.name}</div>
                                </div>
                            </div>
                        </td>
                        <td className="px-4 py-3 whitespace-nowrap text-sm text-gray-500">{booking.source}</td>
                        <td className="px-4 py-3 whitespace-nowrap text-sm text-gray-500">{booking.date}</td>
                        <td className="px-4 py-3 whitespace-nowrap text-sm text-gray-500">{booking.email}</td>
                        <td className="px-4 py-3 whitespace-nowrap text-sm text-gray-500">{booking.mobile}</td>
                        <td className="px-4 py-3 whitespace-nowrap text-sm font-medium text-gray-900">{booking.amount}</td>
                        <td className="px-4 py-3 whitespace-nowrap">
                <span
                    className={`inline-flex rounded-full px-3 py-1 text-xs font-medium ${
                        booking.status === "Cash - Paid"
                            ? "bg-green-100 text-green-800"
                            : booking.status === "Online - Paid"
                                ? "bg-green-100 text-green-800"
                                : "bg-yellow-100 text-yellow-800"
                    }`}
                >
                  {booking.status}
                </span>
                        </td>
                        <td className="px-4 py-3 whitespace-nowrap text-right text-sm font-medium">
                            <button className="text-gray-400 hover:text-gray-500">
                                <MoreHorizontal className="h-5 w-5" />
                            </button>
                        </td>
                    </tr>
                ))}
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
    )
}