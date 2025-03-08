"use client"

import * as React from "react"
import { useState, useRef, type KeyboardEvent } from "react"
import { X, Check, ChevronsUpDown, Plus } from "lucide-react"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Command, CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from "@/components/ui/command"
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover"
import { Badge } from "@/components/ui/badge"

type Option = {
  value: string
  label: string
}

type ComponentProps = {
  BodyClass?: string
}

export default function MultiSelect({ BodyClass = "" }: ComponentProps) {
  const initialOptions: Option[] = [
    { value: "react", label: "React" },
    { value: "nextjs", label: "Next.js" },
    { value: "tailwind", label: "Tailwind CSS" },
    { value: "typescript", label: "TypeScript" },
    { value: "node", label: "Node.js" },
  ]

  const [open, setOpen] = useState(false)
  const [options, setOptions] = useState<Option[]>(initialOptions)
  const [selectedValues, setSelectedValues] = useState<Option[]>([])
  const [inputValue, setInputValue] = useState("")
  const inputRef = useRef<HTMLInputElement>(null)

  const handleSelect = (option: Option) => {
    setSelectedValues((prev) => {
      if (prev.some((item) => item.value === option.value)) {
        return prev.filter((item) => item.value !== option.value)
      } else {
        return [...prev, option]
      }
    })
  }

  const handleRemove = (option: Option) => {
    setSelectedValues((prev) => prev.filter((item) => item.value !== option.value))
  }

  const handleKeyDown = (e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter" && inputValue.trim() !== "") {
      e.preventDefault()

      // Check if option already exists
      if (!options.some((option) => option.value === inputValue.toLowerCase().replace(/\s+/g, "-"))) {
        const newOption = {
          value: inputValue.toLowerCase().replace(/\s+/g, "-"),
          label: inputValue.trim(),
        }

        setOptions((prev) => [...prev, newOption])
        setSelectedValues((prev) => [...prev, newOption])
        setInputValue("")
      }
    }
  }

  return (
    <div className="w-full mx-auto">
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            className={cn("w-full justify-between min-h-[44px] h-auto", BodyClass)}
          >
            <div className="flex flex-wrap gap-1 items-center">
              {selectedValues.length > 0 ? (
                selectedValues.map((option) => (
                  <Badge key={option.value} variant="secondary" className="mr-1 px-2 py-0.5">
                    {option.label}
                    <span
                      type="button"
                      className="hover:pointer ml-1 ring-offset-background rounded-full outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2"
                      onKeyDown={(e) => {
                        if (e.key === "Enter") {
                          e.preventDefault()
                          handleRemove(option)
                        }
                      }}
                      onMouseDown={(e) => {
                        e.preventDefault()
                        e.stopPropagation()
                      }}
                      onClick={(e) => {
                        e.preventDefault()
                        e.stopPropagation()
                        handleRemove(option)
                      }}
                    >
                      <X className="h-3 w-3 text-muted-foreground hover:text-foreground" />
                    </span>
                  </Badge>
                ))
              ) : (
                <span className="text-muted-foreground">Select items...</span>
              )}
            </div>
            <ChevronsUpDown className="h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-full p-0">
          <Command className="w-full bg-white text-black">
            <CommandInput
              placeholder="Search or add new item..."
              value={inputValue}
              onValueChange={setInputValue}
              onKeyDown={handleKeyDown}
              ref={inputRef}
              className="h-9"
            />
            {inputValue && !options.some((option) => option.label.toLowerCase() === inputValue.toLowerCase()) && (
              <CommandItem
                value={`add-${inputValue}`}
                className="text-sm"
                onSelect={() => {
                  const newOption = {
                    value: inputValue.toLowerCase().replace(/\s+/g, "-"),
                    label: inputValue.trim(),
                  }
                  setOptions((prev) => [...prev, newOption])
                  setSelectedValues((prev) => [...prev, newOption])
                  setInputValue("")
                }}
              >
                <Plus className="mr-2 h-4 w-4" />
                Add "{inputValue}"
              </CommandItem>
            )}
            <CommandList>
              <React.Fragment>
                <CommandEmpty>No results found.</CommandEmpty>
                <CommandGroup className="max-h-64 overflow-auto bg-white text-black">
                  {options.map((option) => (
                    <CommandItem key={option.value} value={option.value} onSelect={() => handleSelect(option)}>
                      <div className="flex items-center w-full">
                        <Check
                          className={cn(
                            "mr-2 h-4 w-4",
                            selectedValues.some((item) => item.value === option.value) ? "opacity-100" : "opacity-0",
                          )}
                        />
                        {option.label}
                      </div>
                    </CommandItem>
                  ))}
                </CommandGroup>
              </React.Fragment>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>
    </div>
  )
}

