import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import DataTable, { type Column } from "../components/blocks/admin/DataTable";
import SongForm from "../components/form/admin/SongForm";
import { Button } from "../components/ui/Button";
import { PlusIcon } from "../components/ui/icons";
import type { Track } from "../types";

const AdminSongsPage = () => {
  const [songs, setSongs] = useState<Track[]>([]);
  const [currentSong, setCurrentSong] = useState<Track | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();
  const { id } = useParams();

  const isFormOpen = !!id;

  const onCancel = () => {
    navigate("/admin/songs");
  };

  const formatDate = (isoDate: string) => {
    return new Date(isoDate).toLocaleDateString("ru-RU", {
      day: "numeric",
      month: "long",
      year: "numeric",
    });
  };

  useEffect(() => {
    const fetchTracks = async () => {
      setIsLoading(true);
      try {
        const res = await fetch("http://localhost:8080/api/songs", {
          method: "GET",
          credentials: "include",
        });
        const data = await res.json();
        if (Array.isArray(data)) {
          setSongs(data);
        }
      } catch (error) {
        console.error("Error fetching songs:", error);
      } finally {
        setIsLoading(false);
      }
    };
    fetchTracks();
  }, []);

  useEffect(() => {
    if (id && id !== "new") {
      const numericId = Number(id);
      const fetchSong = async () => {
        try {
          const res = await fetch(
            `http://localhost:8080/api/songs/${numericId}`,
            {
              credentials: "include",
            },
          );
          const data = await res.json();
          setCurrentSong(data);
        } catch (error) {
          console.error("Error fetching song:", error);
        }
      };
      fetchSong();
    } else {
      setCurrentSong(null);
    }
  }, [id]);

  const handleEdit = (song: Track) => {
    navigate(`/admin/songs/${song.id}`);
  };

  const handleDelete = async (id: number) => {
    try {
      const res = await fetch(`http://localhost:8080/api/admin/songs/${id}`, {
        method: "DELETE",
        credentials: "include",
      });
      if (!res.ok) throw new Error("Failed to delete song");
      setSongs((prev) => prev.filter((song) => song.id !== id));
    } catch (err) {
      console.error("Error deleting song:", err);
    }
  };

  const handleSubmit = async (songData: Track, imageFile?: File) => {
    const formData = new FormData();

    if (currentSong) {
      if (!imageFile && currentSong.image_url) {
        formData.append("current_image", currentSong.image_url);
      }
      if (songData.title !== currentSong.title) {
        formData.append("title", songData.title);
      }
      if (songData.full_title !== currentSong.full_title) {
        formData.append("full_title", songData.full_title);
      }
      if (imageFile) {
        formData.append("image", imageFile);
      }
      if (songData.release_date !== currentSong.release_date) {
        formData.append("release_date", songData.release_date);
      }
      if (songData.genre.id !== currentSong.genre.id) {
        formData.append("genre_id", songData.genre.id.toString());
      }

      if (songData.artist.id !== currentSong.artist.id) {
        formData.append("artist_id", songData.artist.id.toString());
      }
      if (songData.album?.id !== currentSong.album?.id) {
        if (songData.album?.id !== undefined && songData.album?.id !== null) {
          formData.append("album_id", songData.album?.id.toString());
        } else {
          formData.append("album_id", "");
        }
      }
    } else {
      formData.append("title", songData.title);
      formData.append("full_title", songData.full_title);
      formData.append("release_date", songData.release_date);
      if (imageFile) {
        formData.append("image", imageFile);
      }
      formData.append("genre_id", songData.genre.id.toString());
      formData.append("artist_id", songData.artist.id.toString());
      if (songData.album?.id !== undefined && songData.album?.id !== null) {
        formData.append("album_id", songData.album?.id.toString());
      }
    }

    try {
      const url = currentSong
        ? `http://localhost:8080/api/admin/songs/${currentSong.id}`
        : `http://localhost:8080/api/admin/songs`;
      const method = currentSong ? "PATCH" : "POST";

      const res = await fetch(url, {
        method,
        credentials: "include",
        body: formData,
      });

      if (!res.ok) {
        const errorData = await res.json();
        throw new Error(errorData.message || "Request failed");
      }

      const updatedSong = await res.json();

      if (currentSong) {
        setSongs((prev) =>
          prev.map((song) => (song.id === currentSong.id ? updatedSong : song)),
        );
      } else {
        const res = await fetch("http://localhost:8080/api/songs", {
          method: "GET",
          credentials: "include",
        });
        const data = await res.json();
        if (Array.isArray(data)) {
          setSongs(data);
        }
      }

      navigate("/admin/songs");
    } catch (err) {
      console.error("Error submitting song:", err);
    }
  };

  const columns: Column<Track>[] = [
    { key: "id", header: "ID" },
    { key: "title", header: "Title" },
    { key: "full_title", header: "FullTitle" },
    {
      key: "image_url",
      header: "Image",
      render: (value, _item) => (
        <img
          src={
            typeof value === "string"
              ? value.startsWith("http")
                ? value
                : `http://localhost:8080/${value}`
              : undefined
          }
          alt="Song"
          className="h-10 w-10 object-cover rounded"
        />
      ),
    },
    {
      key: "release_date",
      header: "ReleaseDate",
      render: (value) => formatDate(value as string),
    },
    {
      key: "genre",
      header: "Genre",
      render: (value, _item) => {
        if (typeof value === "object" && value !== null && "title" in value) {
          return value.title;
        } else {
          return "N/A";
        }
      },
    },
    {
      key: "artist",
      header: "Artist",
      render: (value, _item) => {
        if (
          typeof value === "object" &&
          value !== null &&
          "nickname" in value
        ) {
          return value.nickname;
        } else {
          return "N/A";
        }
      },
    },
    {
      key: "album",
      header: "Album",
      render: (value, _item) => {
        if (typeof value === "object" && value !== null && "title" in value) {
          return value.title;
        } else {
          return "N/A";
        }
      },
    },
  ];

  if (isLoading) return <div>Loading...</div>;

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-gray-800">Manage Songs</h2>
        <Button onClick={() => navigate("/admin/songs/new")}>
          <PlusIcon className="w-5 h-5 mr-1" />
          Добавить песню
        </Button>
      </div>

      {isFormOpen ? (
        <div className="bg-white p-6 rounded-lg shadow-sm mb-6">
          <SongForm
            initialData={currentSong}
            onSubmit={handleSubmit}
            onCancel={onCancel}
          />
        </div>
      ) : (
        <DataTable<Track>
          columns={columns}
          data={songs}
          onEdit={handleEdit}
          onDelete={handleDelete}
        />
      )}
    </div>
  );
};

export default AdminSongsPage;
