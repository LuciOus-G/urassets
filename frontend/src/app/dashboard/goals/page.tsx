"use client"

import { useState, useRef, useEffect } from "react"
import { motion } from "framer-motion"
import { MoreHorizontal, Share, Smile } from "lucide-react"

// Define the card data type
interface CardData {
    id: string
    title: string
    percentage: string
    message: string
    description: string
    priority: number
    gridPosition: number // Single number representing position in the grid
}

// Define a grid position type
interface GridPosition {
    id: number
    row: number
    col: number
    x: number
    y: number
    width: number
    height: number
    isOccupied: boolean
    cardId?: string
}

export default function PriorityCards() {
    const containerRef = useRef<HTMLDivElement>(null)
    const [containerWidth, setContainerWidth] = useState(0)

    // Grid configuration
    const gridGap = 20
    const minCardWidth = 300
    const cardHeight = 350 // Approximate height of a card

    // Calculate columns based on container width
    const columns = Math.max(1, Math.floor((containerWidth + gridGap) / (minCardWidth + gridGap)))
    const gridCellWidth = columns > 0 ? (containerWidth - gridGap * (columns - 1)) / columns : minCardWidth

    // Card is slightly smaller than grid cell to create visual padding
    const cardPadding = 16
    const cardWidth = gridCellWidth - cardPadding * 2

    // Number of rows in the grid
    const rows = 2

    // Initial card data
    const [cards, setCards] = useState<CardData[]>([
        {
            id: "1",
            title: "Pregnancy Goals",
            percentage: "+56%",
            message: "This week your numbers have skyrocketed. Keep it up!",
            description:
                "You are successfully moving towards your goal, your performance is growing, track and share your progress with us.",
            priority: 1,
            gridPosition: 0, // Top-left position
        },
        {
            id: "2",
            title: "Pregnancy Goals",
            percentage: "+56%",
            message: "This week your numbers have skyrocketed. Keep it up!",
            description:
                "You are successfully moving towards your goal, your performance is growing, track and share your progress with us.",
            priority: 2,
            gridPosition: 1, // Top-center position
        },
        {
            id: "3",
            title: "Pregnancy Goals",
            percentage: "+56%",
            message: "This week your numbers have skyrocketed. Keep it up!",
            description:
                "You are successfully moving towards your goal, your performance is growing, track and share your progress with us.",
            priority: 3,
            gridPosition: 2, // Top-right position
        },
    ])

    // Generate grid positions
    const [gridPositions, setGridPositions] = useState<GridPosition[]>([])

    // Track which grid cell is being hovered over during drag
    const [hoveredCell, setHoveredCell] = useState<number | null>(null)

    // Track if a card is currently being dragged
    const [isDragging, setIsDragging] = useState(false)

    // Update container width and grid positions on mount and resize
    useEffect(() => {
        const updateWidth = () => {
            if (containerRef.current) {
                setContainerWidth(containerRef.current.offsetWidth)
            }
        }

        updateWidth()
        window.addEventListener("resize", updateWidth)
        return () => window.removeEventListener("resize", updateWidth)
    }, [])

    // Update grid positions when container width or cards change
    useEffect(() => {
        if (containerWidth === 0) return

        const positions: GridPosition[] = []

        // Create grid positions
        for (let row = 0; row < rows; row++) {
            for (let col = 0; col < columns; col++) {
                const positionId = row * columns + col
                const x = col * (gridCellWidth + gridGap)
                const y = row * (cardHeight + gridGap)

                // Check if this position is occupied by a card
                const occupyingCard = cards.find((card) => card.gridPosition === positionId)

                positions.push({
                    id: positionId,
                    row,
                    col,
                    x,
                    y,
                    width: gridCellWidth,
                    height: cardHeight,
                    isOccupied: !!occupyingCard,
                    cardId: occupyingCard?.id,
                })
            }
        }

        setGridPositions(positions)
    }, [containerWidth, cards, columns, rows, gridCellWidth, gridGap, cardHeight])

    // Calculate priority based on grid position (top-left has highest priority)
    const calculatePriority = (updatedCards: CardData[]): CardData[] => {
        // Sort cards by grid position (which already represents row-major order)
        const sortedCards = [...updatedCards].sort((a, b) => a.gridPosition - b.gridPosition)

        // Assign priorities based on the sorted order
        return sortedCards.map((card, index) => ({
            ...card,
            priority: index + 1,
        }))
    }

    // Find the grid position that contains the given coordinates
    const findGridPositionAtPoint = (x: number, y: number): GridPosition | null => {
        if (gridPositions.length === 0) return null

        // Find the grid position that contains the point
        return (
            gridPositions.find((position) => {
                return (
                    x >= position.x && x <= position.x + position.width && y >= position.y && y <= position.y + position.height
                )
            }) || null
        )
    }

    // Handle drag start
    const handleDragStart = () => {
        setIsDragging(true)
    }

    // Handle drag
    const handleDrag = (id: string, dragX: number, dragY: number) => {
        // Find the grid position at the current drag point
        const positionAtPoint = findGridPositionAtPoint(dragX, dragY)

        if (positionAtPoint) {
            setHoveredCell(positionAtPoint.id)
        } else {
            setHoveredCell(null)
        }
    }

    // Handle drag end
    const handleDragEnd = (id: string, dragX: number, dragY: number) => {
        setIsDragging(false)
        setHoveredCell(null)

        // Find the current position of the card
        const card = cards.find((c) => c.id === id)
        if (!card) return

        // Find the grid position at the drop point
        const positionAtPoint = findGridPositionAtPoint(dragX, dragY)
        if (!positionAtPoint) return

        setCards((prevCards) => {
            // Check if the position is already occupied by another card
            const occupyingCard = prevCards.find((c) => c.id !== id && c.gridPosition === positionAtPoint.id)

            let updatedCards = [...prevCards]

            if (occupyingCard) {
                // Swap positions with the occupying card
                updatedCards = updatedCards.map((c) => {
                    if (c.id === id) {
                        return { ...c, gridPosition: positionAtPoint.id }
                    }
                    if (c.id === occupyingCard.id) {
                        return { ...c, gridPosition: card.gridPosition }
                    }
                    return c
                })
            } else {
                // Just update the position of the current card
                updatedCards = updatedCards.map((c) => {
                    if (c.id === id) {
                        return { ...c, gridPosition: positionAtPoint.id }
                    }
                    return c
                })
            }

            // Recalculate priorities based on new positions
            return calculatePriority(updatedCards)
        })
    }

    // Get the position of a card based on its grid position
    const getCardPosition = (gridPosition: number) => {
        const position = gridPositions.find((pos) => pos.id === gridPosition)
        if (!position) return { x: 0, y: 0 }

        // Center the card within the grid cell
        const centeredX = position.x + cardPadding
        const centeredY = position.y + cardPadding

        return { x: centeredX, y: centeredY }
    }

    return (
        <div className="w-full">
            <div
                ref={containerRef}
                className="relative bg-gray-100 rounded-lg p-4"
                style={{
                    touchAction: "none",
                    minHeight: `${rows * cardHeight + (rows - 1) * gridGap + 40}px`, // Add padding
                }}
            >
                {/* Render grid slots */}
                {gridPositions.map((position) => (
                    <div
                        key={`grid-${position.id}`}
                        className={`absolute border-2 rounded-lg transition-all duration-200 ${
                            position.id === hoveredCell
                                ? "border-blue-500 bg-blue-100/30 border-solid"
                                : position.isOccupied
                                    ? "border-blue-300 border-dashed bg-blue-50/20"
                                    : "border-gray-300 border-dashed"
                        }`}
                        style={{
                            left: position.x,
                            top: position.y,
                            width: position.width,
                            height: position.height,
                        }}
                    />
                ))}

                {/* Render cards */}
                {cards.map((card) => {
                    const position = getCardPosition(card.gridPosition)

                    return (
                        <motion.div
                            key={card.id}
                            className="absolute cursor-grab active:cursor-grabbing touch-none"
                            style={{
                                width: cardWidth,
                                height: cardHeight - cardPadding * 2,
                                zIndex: isDragging ? 1000 : 900 - card.priority, // Higher priority cards appear on top
                            }}
                            initial={{ x: position.x, y: position.y }}
                            animate={{
                                x: position.x,
                                y: position.y,
                                transition: {
                                    type: "spring",
                                    damping: 20,
                                    stiffness: 300,
                                },
                            }}
                            drag
                            dragMomentum={false}
                            onDragStart={handleDragStart}
                            onDrag={(_, info) => {
                                const dragX = info.point.x
                                const dragY = info.point.y
                                handleDrag(card.id, dragX, dragY)
                            }}
                            onDragEnd={(_, info) => {
                                const dragX = info.point.x
                                const dragY = info.point.y
                                handleDragEnd(card.id, dragX, dragY)
                            }}
                            whileDrag={{
                                scale: 1.03,
                                boxShadow: "0 10px 25px rgba(0, 0, 0, 0.15)",
                                zIndex: 1000,
                            }}
                        >
                            <div className="bg-blue-500 text-white rounded-lg p-6 shadow-lg h-full flex flex-col">
                                <div className="flex justify-between items-center mb-6">
                                    <div className="flex items-center gap-2">
                                        <div className="bg-white rounded-full p-1">
                                            <Smile className="h-6 w-6 text-blue-500" />
                                        </div>
                                        <span className="font-medium">{card.title}</span>
                                    </div>
                                    <div className="flex items-center gap-2">
                    <span className="bg-white text-blue-500 rounded-full px-2 py-1 text-xs font-bold">
                      Priority: {card.priority}
                    </span>
                                        <button className="text-white hover:bg-blue-600 rounded-full p-1">
                                            <MoreHorizontal className="h-5 w-5" />
                                        </button>
                                    </div>
                                </div>

                                <div className="flex-grow">
                                    <div className="text-6xl font-bold mb-4">{card.percentage}</div>
                                    <p className="text-lg mb-4">{card.message}</p>
                                    <p className="text-sm opacity-90 mb-6">{card.description}</p>
                                </div>

                                <button className="bg-white text-blue-500 rounded-full py-2 px-6 flex items-center justify-center gap-2 w-max hover:bg-blue-50 transition-colors">
                                    <Share className="h-4 w-4" />
                                    <span>Share</span>
                                </button>
                            </div>
                        </motion.div>
                    )
                })}
            </div>

            <div className="mt-8 p-4 bg-white rounded-lg shadow">
                <h2 className="text-lg font-semibold mb-2">Current Priority Order:</h2>
                <ul className="list-disc pl-5">
                    {[...cards]
                        .sort((a, b) => a.priority - b.priority)
                        .map((card) => {
                            const position = gridPositions.find((pos) => pos.id === card.gridPosition)
                            return (
                                <li key={card.id} className="mb-1">
                                    Card {card.id}: Priority {card.priority} (Position: Row {(position?.row ?? 0) + 1}, Column{" "}
                                    {(position?.col ?? 0) + 1})
                                </li>
                            )
                        })}
                </ul>
                <p className="mt-4 text-sm text-gray-500">
                    Drag cards to any grid slot. Cards will automatically center within the slot. Priority is determined by
                    position (top to bottom, left to right).
                </p>
            </div>
        </div>
    )
}