import { useEffect, useState } from "react";
import type { Track } from "../types";

export const useTracksRating = () => {
  const [tracks, setTracks] = useState<Track[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchTracksSorted = async () => {
      try {
        const res = await fetch(
          "http://localhost:8080/api/songs/sorted-by-rating",
        );
        const data: Track[] = await res.json();
        setTracks(data);
      } catch (e) {
        console.log("Ошибка загрузки рейтинга", e);
      } finally {
        setLoading(false);
      }
    };
    fetchTracksSorted();
  }, []);

  return { tracks, loading };
};
