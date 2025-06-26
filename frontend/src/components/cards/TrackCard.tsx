import type { FC } from "react";
import type { Track } from "../../types";
import { LikeIcon, DislikeIcon } from "../ui/icons";

const TrackCard: FC<{ track: Track; onClick: (track: Track) => void }> = ({
  track,
  onClick,
}) => {
  const formatDate = (isoDate: string) => {
    return new Date(isoDate).toLocaleDateString("ru-RU", {
      day: "numeric",
      month: "short",
      year: "numeric",
    });
  };

  return (
    <div
      onClick={() => onClick(track)}
      className="w-full flex flex-col gap-3 cursor-pointer group relative"
    >
      <div className="relative rounded-xl overflow-hidden shadow-lg aspect-square">
        <img
          src={"http://localhost:8080/" + track.image_url}
          alt={track.full_title}
          className="w-full h-full object-cover transition-all duration-500 
                     group-hover:scale-105 brightness-90 group-hover:brightness-50"
        />

        {track.genre && (
          <span className="absolute top-2 left-2 bg-gray-900/80 text-xs text-white px-2 py-1 rounded-full">
            {track.genre.title}
          </span>
        )}

        <div className="absolute bottom-2 right-2 flex items-center gap-2 bg-gray-900/80 backdrop-blur-sm rounded-full px-3 py-1">
          <div className="flex items-center gap-1 text-green-500">
            <LikeIcon className="w-3 h-3" />
            <span className="text-xs font-medium">{track.like_count}</span>
          </div>
          <div className="w-px h-4 bg-gray-600"></div>
          <div className="flex items-center gap-1 text-red-500">
            <DislikeIcon className="w-3 h-3" />
            <span className="text-xs font-medium">{track.dislike_count}</span>
          </div>
        </div>
      </div>

      <div className="flex flex-col gap-1 px-1">
        <h3 className="text-lg font-bold text-white truncate hover:text-purple-400 transition-colors">
          {track.title}
        </h3>
        <div className="flex justify-between items-center">
          <p className="text-sm text-gray-400 truncate">
            {track.artist.nickname}
          </p>
          <p className="text-xs text-gray-500 whitespace-nowrap ml-2">
            {formatDate(track.release_date)}
          </p>
        </div>
      </div>
    </div>
  );
};

export { TrackCard };
