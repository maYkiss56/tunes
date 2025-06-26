import { useEffect, useState } from "react";
import type { Track } from "../../types";
import { TrackCard } from "../cards/TrackCard";
import ModelWindow from "../blocks/ModelWindow";

const TrackList = () => {
  const [tracks, setTracks] = useState<Track[]>([]);
  const [loading, setLoading] = useState(true);
  const [selectedTrack, setSelectedTrack] = useState<Track | null>(null);
  useEffect(() => {
    const fetchTracks = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/songs");
        if (!res.ok) throw new Error("Ошибка при загрузке песен");
        const data = await res.json();
        setTracks(data);
      } catch (error) {
        console.log("Ошибка при получении песен: ", error);
      } finally {
        setLoading(false);
      }
    };

    fetchTracks();
  }, []);

  if (loading)
    return <p className="text-center text-white">Загрузка песен...</p>;

  return (
    <>
      <section className="w-304 grid items-center justify-between mx-auto grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-6 p-4">
        {tracks.map((track) => (
          <TrackCard key={track.id} track={track} onClick={setSelectedTrack} />
        ))}
      </section>

      {selectedTrack && (
        <ModelWindow
          track={selectedTrack}
          onClose={() => setSelectedTrack(null)}
        />
      )}
    </>
  );
};

export { TrackList };
