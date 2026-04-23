import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { endpoints, type EndpointStatus } from "../../endpoints";

async function checkEndpoint(port: number, endpoint: string): Promise<boolean> {
  try {
    const res = await fetch(`http://localhost:${port}${endpoint}`);
    return res.ok;
  } catch {
    return false;
  }
}

function Status() {
  const [results, setResults] = useState<EndpointStatus[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const runChecks = async () => {
      const checks = await Promise.all(
        endpoints.map(async (ep) => ({
          ...ep,
          ok: await checkEndpoint(ep.port, ep.endpoint),
        }))
      );
      setResults(checks);
      setLoading(false);
    };

    runChecks();
  }, []);

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center text-yellow-400">
        Checking server status...
      </div>
    );
  }

  return (
    <div className="mt-8 min-h-screen flex flex-col gap-4 items-center font-mono min-w-full">
      <h1 className="text-3xl font-bold mb-2">Server Status</h1>
      {results.map((r) => (
        <div
          key={r.endpoint}
          className="md:min-w-md min-w-sm flex items-center justify-between w-80 bg-gray-900 px-4 py-3 rounded-lg"
        >
          <span>
            <Link
              to={`http://${r.addr}:${r.port}${r.endpoint}`}
              className="group text-green-300 font-mono relative inline-block ml-2"
            >
              {r.name}
              <span className="ml-2 opacity-0 group-hover:opacity-100 transition text-xs text-green-400">
                Click
              </span>
            </Link>
          </span>
          <span
            className={r.ok ? "text-green-400" : "text-red-400"}>
            {r.ok ? "OK" : "KO"}
          </span>
        </div>
      ))}
    </div>
  );
}

export default Status;