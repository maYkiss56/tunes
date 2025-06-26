import { useState, useEffect } from "react";
import type { Track } from "../types";

export const useTopTracks = (
  timeRange: string = "week",
  limit: number = 20,
) => {
  const [tracks, setTracks] = useState<Track[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchTopTracks = async () => {
      try {
        setLoading(true);
        const url = `http://localhost:8080/api/songs/top?time_range=${timeRange}&limit=${limit}`;
        const response = await fetch(url);

        if (!response.ok) {
          throw new Error("Ошибка при загрузке топовых треков");
        }

        const data = await response.json();
        setTracks(data);
      } catch (error) {
        console.error("Ошибка при загрузке топовых треков:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchTopTracks();
  }, [timeRange, limit]);

  return { tracks, loading };
};
