"use client"
import {Plus, ChevronDown, Pencil} from "lucide-react";
import Modal from "@/components/modal";
import {Dispatch, SetStateAction, useState} from "react";

type bankList = {
  id: string
  name: string
  icon: string
}

type expensesCategory = {
  id: string
  name: string
}

type transferFeeCategory = {
  id: number
  name: string
  amount: number
}

const bankList: bankList[] = [
  { id: "bca", name: "Bank Central Asia", icon: "https://www.bca.co.id/-/media/Feature/Card/List-Card/Tentang-BCA/Brand-Assets/Logo-BCA/Logo-BCA_Biru.png" },
  { id: "cimb", name: "CIMB Niaga", icon: "https://web14.bernama.com/storage/photos/49b413c6b4bd818fc88f6cb9648add1d63ef7b048681f" },
  { id: "jago", name: "Bank Jago",  icon: "https://www.jago.com/images/media-assets/combi1.svg" }
]

const expensesCategory: expensesCategory[] = [
  { id: "food", name: "Food" },
  { id: "transportation", name: "Transportation" },
  { id: "tw", name: "Transfer Wealth" }
]

const transferFeeCategories: transferFeeCategory[] = [
  { id: 1, name: "No", amount: 0 },
  { id: 2, name: "Bi Fast Rp. 2,500", amount: 2500 },
  { id: 3, name: "RTGS Rp. 30,000", amount: 30000 },
  { id: 4, name: "RTO Rp. 4,000", amount: 4000 },
  { id: 5, name: "Regular Rp. 6,500", amount: 6500 },
  { id: 6, name: "Custom", amount: 0 }
]

export default function Page() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isMutation, setIsMutation] = useState(false);
  const [formData, setFormData] = useState({
    fromBankId: '',
    toBankId: '',
    amount: 0,
    categoryId: '',
    transferFee: 0,
    descriptions: ''
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
  const [fromBankOpen, setFromBankOpen] = useState(false)
  const [toBankOpen, setToBankOpen] = useState(false)
  const [transferFeeOpen, setTransferFeeOpen] = useState(false)
  const [customTransferFeeOpen, setCustomTransferFeeOpen] = useState(false)


  const [selectedToBankCategory, setSelectedToBankCategory] = useState(bankList[0])
  const [selectedFromBankCategory, setSelectedFromBankCategory] = useState(bankList[0])
  const [selectedExpensesCategory, setSelectedExpensesCategory] = useState(expensesCategory[0])
  const [selectedTransferFeeCategory, setSelectedTransferFeeCategory] = useState(transferFeeCategories[0])

  const toggleDropdown = (state: boolean, setState: Dispatch<SetStateAction<boolean>>) => {
    setState(!state)
  }


  const handleCategoryChange = (selectedBank: bankList, setBankChange: Dispatch<SetStateAction<any>>, setBankDropDown: Dispatch<SetStateAction<boolean>>) => {
    setBankChange(selectedBank)
    setBankDropDown(false)
  }

  const handleTransferFeeCategory = (tf: transferFeeCategory) => {
    if (tf.id === 6) {
      setCustomTransferFeeOpen(true)
    } else {
      setCustomTransferFeeOpen(false)
    }
    setSelectedTransferFeeCategory(tf)
    setTransferFeeOpen(false)
  }

  const handleExpensesCategory = (ec: expensesCategory) => {
    if (ec.id === 'tw') {
      setIsMutation(true)
    } else {
      setIsMutation(false)
    }
    setSelectedExpensesCategory(ec)
    setIsOpen(false)
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>, isNumber: boolean) => {
    let val = e.target.value
    if (isNumber) {
      val = formatRupiah(e.target.value.replace(/[^0-9]/g, ""));
    }
    setFormData({ ...formData, [e.target.name]: val });
  };

  return (
    <div>
      <div className="flex justify-end items-center">
        <button onClick={() => setIsModalOpen(true)} className="flex justify-beetwen items-center bg-green-400 p-2 rounded-lg">
          <Plus className="text-green-900"/>
          <h1 className="text-sm font-bold text-green-900">Add Transactions</h1>
        </button>

        <Modal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)}>
          <h1 className="text-2xl text-black text-bold">Add Transactions</h1>
          <form>
            <div className="w-full max-w-md mx-auto py-3">
              <div className="flex items-center gap-2 mb-2">
                <label htmlFor="expensesCategory" className="text-base font-bold text-black">Select category</label>
              </div>

              <div className="relative">
                <button
                  type="button"
                  onClick={() => toggleDropdown(isOpen, setIsOpen)}
                  className="flex items-center justify-between w-full p-2 bg-white border border-gray-200 rounded-full shadow-sm"
                  aria-haspopup="listbox"
                  aria-expanded={isOpen}
                >
                  <div className="flex items-center gap-3">
                    <span className="text-base text-black">{selectedExpensesCategory.name}</span>
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
                      {expensesCategory.map((ec) => (
                        <div key={ec.id} className="relative">
                          <input
                            type="radio"
                            id={`expensesCategory-${ec.id}`}
                            name="categoryId"
                            value={ec.id}
                            checked={selectedExpensesCategory.id === ec.id}
                            onChange={(e) => {
                              handleExpensesCategory(ec)
                              handleChange(e, false)
                            }}
                            className="absolute opacity-0 w-full h-full cursor-pointer"
                          />
                          <label
                            htmlFor={`expensesCategory-${ec.id}`}
                            className={`flex items-center gap-3 p-4 cursor-pointer hover:bg-gray-50 ${
                              selectedExpensesCategory.id === ec.id ? "bg-gray-50" : ""
                            }`}
                          >
                            <span className="text-base text-black">{ec.name}</span>
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
                <label htmlFor="fromBank" className="text-base font-bold text-black">From Bank</label>
              </div>

              <div className="relative">
                <button
                  type="button"
                  onClick={() => toggleDropdown(fromBankOpen, setFromBankOpen)}
                  className="flex items-center justify-between w-full p-2 bg-white border border-gray-200 rounded-full shadow-sm"
                  aria-haspopup="listbox"
                  aria-expanded={fromBankOpen}
                >
                  <div className="flex items-center gap-3">
                        <span className="text-base">
                          <img width={50} height={50} alt={selectedFromBankCategory.name} src={selectedFromBankCategory.icon} />
                        </span>
                    <span className="text-base text-black">{selectedFromBankCategory.name}</span>
                  </div>
                  <div
                    className={`w-10 h-10 rounded-full bg-gray-100 flex items-center justify-center transition-transform ${
                      fromBankOpen ? "rotate-180" : ""
                    }`}
                  >
                    <ChevronDown className="w-5 h-5 text-gray-500" />
                  </div>
                </button>

                {fromBankOpen && (
                  <div className="absolute z-10 w-full mt-1 bg-white border border-gray-200 rounded-2xl shadow-lg max-h-60 overflow-auto">
                    <fieldset className="p-1">
                      <legend className="sr-only">Select a category</legend>
                      {bankList.map((bankItem) => (
                        <div key={bankItem.id} className="relative">
                          <input
                            type="radio"
                            id={`fromBank-${bankItem.id}`}
                            name="fromBankId"
                            value={bankItem.id}
                            checked={selectedFromBankCategory.id === bankItem.id}
                            onChange={(e) => {
                              handleCategoryChange(bankItem, setSelectedFromBankCategory, setFromBankOpen);
                              handleChange(e, false);
                            }}
                            className="absolute opacity-0 w-full h-full cursor-pointer"
                          />
                          <label
                            htmlFor={`fromBank-${bankItem.id}`}
                            className={`flex items-center gap-3 p-4 cursor-pointer hover:bg-gray-50 ${
                              selectedFromBankCategory.id === bankItem.id ? "bg-gray-50" : ""
                            }`}
                          >
                                <span className="text-base text-black">
                                  <img width={50} height={50} alt={bankItem.name} src={bankItem.icon} />
                                </span>
                            <span className="text-base text-black">{bankItem.name}</span>
                          </label>
                        </div>
                      ))}
                    </fieldset>
                  </div>
                )}
              </div>
            </div>

            {isMutation && (
              <>
                <div className="w-full max-w-md mx-auto py-3">
                  <div className="flex items-center gap-2 mb-2">
                    <label htmlFor="toBank" className="text-base font-bold text-black">To Bank</label>
                  </div>

                  <div className="relative">
                    <button
                      type="button"
                      onClick={() => toggleDropdown(toBankOpen, setToBankOpen)}
                      className="flex items-center justify-between w-full p-2 bg-white border border-gray-200 rounded-full shadow-sm"
                      aria-haspopup="listbox"
                      aria-expanded={toBankOpen}
                    >
                      <div className="flex items-center gap-3">
                        <span className="text-base">
                          <img width={50} height={50} alt={selectedToBankCategory.name} src={selectedToBankCategory.icon} />
                        </span>
                        <span className="text-base text-black">{selectedToBankCategory.name}</span>
                      </div>
                      <div
                        className={`w-10 h-10 rounded-full bg-gray-100 flex items-center justify-center transition-transform ${
                          toBankOpen ? "rotate-180" : ""
                        }`}
                      >
                        <ChevronDown className="w-5 h-5 text-gray-500" />
                      </div>
                    </button>

                    {toBankOpen && (
                      <div className="absolute z-10 w-full mt-1 bg-white border border-gray-200 rounded-2xl shadow-lg max-h-60 overflow-auto">
                        <fieldset className="p-1">
                          <legend className="sr-only">Select a category</legend>
                          {bankList.map((bankItem) => (
                            <div key={bankItem.id} className="relative">
                              <input
                                type="radio"
                                id={`toBank-${bankItem.id}`}
                                name="toBankId"
                                value={bankItem.id}
                                checked={selectedToBankCategory.id === bankItem.id}
                                onChange={(e) => {
                                  handleCategoryChange(bankItem, setSelectedToBankCategory, setToBankOpen);
                                  handleChange(e, false);
                                }}
                                className="absolute opacity-0 w-full h-full cursor-pointer"
                              />
                              <label
                                htmlFor={`toBank-${bankItem.id}`}
                                className={`flex items-center gap-3 p-4 cursor-pointer hover:bg-gray-50 ${
                                  selectedToBankCategory.id === bankItem.id ? "bg-gray-50" : ""
                                }`}
                              >
                                <span className="text-base text-black">
                                  <img width={50} height={50} alt={bankItem.name} src={bankItem.icon} />
                                </span>
                                <span className="text-base text-black">{bankItem.name}</span>
                              </label>
                            </div>
                          ))}
                        </fieldset>
                      </div>
                    )}
                  </div>
                </div>
              </>
            )}

            <div className="w-full max-w-md mx-auto py-3">
              <div className="flex items-center gap-2 mb-2">
                <label htmlFor="isTransferFee" className="text-base font-bold text-black">Transfer Fee</label>
              </div>

              <div className="relative">
                <button
                  type="button"
                  onClick={() => toggleDropdown(transferFeeOpen, setTransferFeeOpen)}
                  className="flex items-center justify-between w-full p-2 bg-white border border-gray-200 rounded-full shadow-sm"
                  aria-haspopup="listbox"
                  aria-expanded={transferFeeOpen}
                >
                  <div className="flex items-center gap-3">
                    <span className="text-base text-black">{selectedTransferFeeCategory.name}</span>
                  </div>
                  <div
                    className={`w-10 h-10 rounded-full bg-gray-100 flex items-center justify-center transition-transform ${isOpen ? "rotate-180" : ""}`}
                  >
                    <ChevronDown className="w-5 h-5 text-gray-500" />
                  </div>
                </button>

                {transferFeeOpen && (
                  <div className="absolute z-10 w-full mt-1 bg-white border border-gray-200 rounded-2xl shadow-lg max-h-60 overflow-auto">
                    <fieldset className="p-1">
                      <legend className="sr-only">Select a category</legend>
                      {transferFeeCategories.map((tf) => (
                        <div key={tf.id} className="relative">
                          <input
                            type="radio"
                            id={`isTransferFee-${tf.id}`}
                            value={tf.amount}
                            name="transferFee"
                            checked={selectedTransferFeeCategory.id === tf.id}
                            onChange={(e) => {
                              handleTransferFeeCategory(tf)
                              handleChange(e, false)
                            }}
                            className="absolute opacity-0 w-full h-full cursor-pointer"
                          />
                          <label
                            htmlFor={`isTransferFee-${tf.id}`}
                            className={`flex items-center gap-3 p-4 cursor-pointer hover:bg-gray-50 ${
                              selectedTransferFeeCategory.id === tf.id ? "bg-gray-50" : ""
                            }`}
                          >
                            <span className="text-base text-black">{tf.name}</span>
                          </label>
                        </div>
                      ))}
                    </fieldset>
                  </div>
                )}
              </div>
            </div>

            {customTransferFeeOpen && (
              <div className="w-full max-w-md mx-auto py-3">
                <div className="flex items-center gap-2 mb-2">
                  <label htmlFor={"transferFee"} className="text-base font-bold text-black">Custom Transfer Fee</label>
                </div>

                <div className="relative">
                  <div
                    type="button"
                    className="flex items-center justify-between w-full p-2 bg-white border border-gray-200 rounded-full shadow-sm">
                    <input
                      className="rounded-full w-full p-2 text-black"
                      name="transferFee"
                      value={formData.transferFee}
                      onChange={(e) => handleChange(e, true)}
                    />
                  </div>
                </div>
              </div>
            )}

            <div className="w-full max-w-md mx-auto py-3">
              <div className="flex items-center gap-2 mb-2">
                <label htmlFor={"amount"} className="text-base font-bold text-black">Amount</label>
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
                    name="amount"
                    value={formData.amount}
                    onChange={(e) => handleChange(e, true)}
                  />
                </div>
              </div>
            </div>

            <div className="w-full max-w-md mx-auto py-3">
              <div className="flex items-center gap-2 mb-2">
                <label htmlFor={"descriptions"} className="text-base font-bold text-black">Descriptions</label>
              </div>

              <div className="relative">
                <div
                  type="button"
                  className="flex items-center justify-between w-full p-2 bg-white border border-gray-200 rounded-full shadow-sm"
                  aria-haspopup="listbox"
                  aria-expanded={isOpen}
                >
                  <input
                    id={"descriptions"}
                    className="rounded-full w-full p-2 text-black"
                    name="descriptions"
                    value={formData.descriptions}
                    onChange={(e) => handleChange(e, false)}
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
        <h1 className="text-black text-xl text-bold">List of Transactions and Mutations</h1>

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

