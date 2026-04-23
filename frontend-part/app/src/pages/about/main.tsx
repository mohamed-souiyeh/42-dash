import { Link } from "react-router-dom";
import { Code2, Cpu, Palette, Zap } from "lucide-react";

type TechItemProps = {
    icon: React.ReactNode;
    label: string;
};

function TechItem({ icon, label }: TechItemProps) {
    return (
        <span className="flex items-center gap-2 text-white font-medium">
            {icon}
            {label}
        </span>
    );
}

function About() {
    return (
        <div className="pb-12 min-h-screen flex flex-col items-center
                justify-center bg-gray-950 text-white px-6 text-center">
            <h1 className="text-4xl/12 font-bold max-w-sm">
                Welcome to the DASH project! (,,&gt;_&lt;,,)
            </h1>

            <p className="mt-8 text-gray-400 max-w-sm text-justify">
                This web application is part of a Dash challenge from 42 Málaga,
                where the goal is to build a frontend that consumes and renders
                data from a backend API.
            </p>

            <div className="mt-6 flex flex-wrap gap-4 justify-center">
                <TechItem icon={<Code2 size={16} />} label="React" />
                <TechItem icon={<Cpu size={16} />} label="TypeScript" />
                <TechItem icon={<Palette size={16} />} label="TailwindCSS" />
                <TechItem icon={<Zap size={16} />} label="Vite" />
            </div>

            <p className="mt-6 text-gray-400 max-w-sm text-justify">
                The project focuses on rendering dynamic game data from a backend
                service, practicing component design, state management, and API integration.
            </p>

            <p className="mt-6 text-sm text-gray-500 flex items-center justify-center gap-2">
                Frontend: iamrani- {" "}
                <a
                    href="https://github.com/abykko"
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex items-center gap-1 text-green-400 hover:text-green-200 transition"
                >
                    @abykko
                </a>
            </p>

            <p className="mt-12 text-gray-400 text-center justify-center items-center">
                <span className="text-gray-300">
                    try to navigate to <Link to="/status" className="group text-green-300">
                        /status
                    </Link>
                </span>
            </p>
        </div>
    );
}

export default About;