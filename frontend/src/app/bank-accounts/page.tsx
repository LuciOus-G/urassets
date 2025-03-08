"use client"
import {Plus, ChevronDown, Pencil} from "lucide-react";
import Modal from "@/components/modal";
import {useState} from "react";

type Category = {
  id: string
  name: string
  icon: string
}

const categories: Category[] = [
  { id: "bca", name: "Bank Central Asia", icon: "https://www.bca.co.id/-/media/Feature/Card/List-Card/Tentang-BCA/Brand-Assets/Logo-BCA/Logo-BCA_Biru.png" },
  { id: "cimb", name: "CIMB Niaga", icon: "https://web14.bernama.com/storage/photos/49b413c6b4bd818fc88f6cb9648add1d63ef7b048681f" },
  { id: "jago", name: "Bank Jago",  icon: "https://www.jago.com/images/media-assets/combi1.svg" }
]

export default function Page() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [formData, setFormData] = useState({
    bankName: '',
    balance: ''
  });
  const formatRupiah = (num: string) => {
    if (!num) return ""; // Handle empty input
    const numericValue = num.replace(/\D/g, ""); // Remove non-numeric characters
    return new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
      minimumFractionDigits: 0,
    }).format(Number(numericValue));
  };


  const [isOpen, setIsOpen] = useState(false)
  const [selectedCategory, setSelectedCategory] = useState(categories[0])
  const toggleDropdown = () => setIsOpen(!isOpen)
  const handleCategoryChange = (category: Category) => {
    setSelectedCategory(category)
    setIsOpen(false)
  }



  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    let val = e.target.value
    if (e.target.name === "balance") {
      val = formatRupiah(e.target.value.replace(/[^0-9]/g, ""));
    }
    setFormData({ ...formData, [e.target.name]: val });
  };



  return (
    <div>
      <div className="flex justify-end items-center">
        <button onClick={() => setIsModalOpen(true)} className="flex justify-beetwen items-center bg-green-400 p-2 rounded-lg">
          <Plus className="text-green-900"/>
          <h1 className="text-sm font-bold text-green-900">Add bank accounts</h1>
        </button>

        <Modal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)}>
          <h1 className="text-2xl text-black text-bold">Add Bank Account</h1>
          <form>
            <div className="w-full max-w-md mx-auto py-3">
              <div className="flex items-center gap-2 mb-2">
                <label htmlFor="bankName" className="text-base font-bold text-black">Select Bank</label>
              </div>

              <div className="relative">
                <button
                  type="button"
                  onClick={toggleDropdown}
                  className="flex items-center justify-between w-full p-2 bg-white border border-gray-200 rounded-full shadow-sm"
                  aria-haspopup="listbox"
                  aria-expanded={isOpen}
                >
                  <div className="flex items-center gap-3">
                    <span className="text-base"><img width={50} height={50} alt={selectedCategory.name} src={selectedCategory.icon}/></span>
                    <span className="text-base text-black">{selectedCategory.name}</span>
                  </div>
                  <div
                    className={`w-10 h-10 rounded-full bg-gray-100 flex items-center justify-center transition-transform ${isOpen ? "rotate-180" : ""}`}
                  >
                    <ChevronDown className="w-5 h-5 text-gray-500" />
                  </div>
                </button>

                {isOpen && (
                  <div className="absolute z-10 w-full mt-1 bg-white border border-gray-200 rounded-2xl shadow-lg max-h-60 overflow-auto">
                    <fieldset className="p-1">
                      <legend className="sr-only">Select a category</legend>
                      {categories.map((category) => (
                        <div key={category.id} className="relative">
                          <input
                            type="radio"
                            id={`category-${category.id}`}
                            name="bankName"
                            value={category.id}
                            checked={selectedCategory.id === category.id}
                            onChange={(e) => {
                              handleCategoryChange(category)
                              handleChange(e)
                            }}
                            className="absolute opacity-0 w-full h-full cursor-pointer"
                          />
                          <label
                            htmlFor={`category-${category.id}`}
                            className={`flex items-center gap-3 p-4 cursor-pointer hover:bg-gray-50 ${
                              selectedCategory.id === category.id ? "bg-gray-50" : ""
                            }`}
                          >
                            <span className="text-base text-black"><img width={50} height={50} alt={category.name} src={category.icon}/></span>
                            <span className="text-base text-black">{category.name}</span>
                          </label>
                        </div>
                      ))}
                    </fieldset>
                  </div>
                )}
              </div>
            </div>

            <div className="w-full max-w-md mx-auto py-3">
              <div className="flex items-center gap-2 mb-2">
                <label htmlFor={"balance"} className="text-base font-bold text-black">Input Initial Balance</label>
              </div>

              <div className="relative">
                <div
                  type="button"
                  className="flex items-center justify-between w-full p-2 bg-white border border-gray-200 rounded-full shadow-sm"
                  aria-haspopup="listbox"
                  aria-expanded={isOpen}
                >
                  <input
                    className="rounded-full w-full p-2 text-black"
                    name="balance"
                    value={formData.balance}
                    onChange={handleChange}
                  />
                </div>
              </div>
            </div>

            <div className="flex justify-end align-center pt-5">
              <button className="bg-green-400 text-green-900 px-4 py-3 rounded-2xl">Save</button>
            </div>
          </form>
        </Modal>
      </div>

      <div className="my-7 p-3 bg-white rounded-lg shadow-lg min-h-[80vh] relative fixed">
        <h1 className="text-black text-xl text-bold">List of bank accounts</h1>

        <div className="shadow-[1px_1px_32px_-10px_rgba(0,0,0,0.4)] p-3 bg-white rounded-lg my-5 overflow-y-auto">
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
                <p className="font-medium text-black">Meta subscription</p>
                <p className="text-sm text-gray-500">07 Sep 2023 at 10:02 AM</p>
              </div>
            </div>

            <div>
              <p className="font-medium text-black">Balance</p>
              <p className="text-sm text-green-500">123123</p>
            </div>

            <button className="p-2 px-4 bg-blue-400 rounded-lg text-black flex justify-beetwen items-center align-center gap-2">
              <Pencil className="w-4 h-4" />
              <span>Edit</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

