export const PORT = 3500;
export const ADDR = "localhost";

export const endpoints = [
  { name: "/healthz", endpoint: "/healthz", addr: ADDR, port: PORT },
  { name: "/api/games", endpoint: "/api/games", addr: ADDR, port: PORT },
  { name: "/api/games?name=cosmic", endpoint: "/api/games?name=cosmic", addr: ADDR, port: PORT },
  { name: "/api/games/:id", endpoint: "/api/games/game-00001-elk-studios", addr: ADDR, port: PORT },
];

export type EndpointStatus = {
  name: string;
  endpoint: string;
  addr: string;
  port: number;
  ok: boolean;
};