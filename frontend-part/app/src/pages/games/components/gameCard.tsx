import { ChevronDown } from "lucide-react";
import type { GameCardProps } from "./gameCardProperties";

function GameCard({
    name = "Mockon",
    provider = "MK",
    category = "SlotMock",
}: GameCardProps) {

    return (
        <div
            className="
                group
                text-green-300
                relative
                inline-block
                min-w-sm
                max-w-sm
                md:min-w-md
                md:max-w-md
                lg:min-w-lg
                min-w-sm
                rounded-xl
                overflow-hidden
                shadow-lg
                border-2"
            data-testid="game-card"
            game-category={category}
            game-provider={provider}
        >
            <div className="px-4 py-2 space-y-1">
                <div className="flex items-center justify-between">
                    <h3 className="text-xl font-semibold text-green-400">
                        {name}
                    </h3>
                    <p className="category text-lg font-light text-gray-500">
                        🕹️ {category}
                    </p>
                </div>
                <div className="flex items-center justify-between gap-2">
                    <p className="provider text-xl font-light text-gray-300">
                        {provider}
                    </p>
                    <button className="text-green-400">
                        <ChevronDown size={35}/>
                    </button>
                </div>
            </div>
        </div>
    );
}

export default GameCard;