import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import DataTable from "../components/blocks/admin/DataTable";
import ArtistForm from "../components/form/admin/ArtistForm";
import { Button } from "../components/ui/Button";
import { PlusIcon } from "../components/ui/icons";
import type { Artist } from "../types";

const AdminArtistsPage = () => {
  const [artists, setArtists] = useState<Artist[]>([]);
  const [currentArtist, setCurrentArtist] = useState<Artist | null>(null);
  const navigate = useNavigate();
  const { id } = useParams();

  const isFormOpen = !!id;

  const onCancel = () => {
    navigate("/admin/artists");
  };

  useEffect(() => {
    fetch("http://localhost:8080/api/artists", {
      method: "GET",
      credentials: "include",
    })
      .then((res) => res.json())
      .then((data: Artist[]) => {
        if (Array.isArray(data)) {
          setArtists(data);
        } else {
          console.error("Data is not array");
        }
      })
      .catch((error) => {
        console.error("Error fetching artists:", error);
      });
  }, []);

  useEffect(() => {
    if (id) {
      const numericId = Number(id);
      const artistToEdit = artists.find((artist) => artist.id === numericId);
      if (artistToEdit) {
        setCurrentArtist(artistToEdit);
      } else if (id === "new") {
        setCurrentArtist(null);
      }
    }
  }, [id, artists]);

  const handleEdit = (artist: Artist) => {
    navigate(`/admin/artists/${artist.id}`);
  };

  const handleDelete = async (id: number) => {
    try {
      const res = await fetch(`http://localhost:8080/api/admin/artists/${id}`, {
        method: "DELETE",
        credentials: "include",
      });

      if (!res.ok) throw new Error("Failed to delete artist");

      setArtists((prev) => prev.filter((artist) => artist.id !== id));
    } catch (err) {
      console.error("Error deleting artist:", err);
    }
  };

  const handleSubmit = async (artistData: Artist) => {
    const payload = {
      nickname: artistData.nickname,
      bio: artistData.bio,
      country: artistData.country,
    };

    try {
      let res: Response;
      let updatedArtist: Artist;

      if (currentArtist) {
        res = await fetch(
          `http://localhost:8080/api/admin/artists/${currentArtist.id}`,
          {
            method: "PATCH",
            credentials: "include",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(payload),
          },
        );

        if (!res.ok) throw new Error("Failed to update artist");
        updatedArtist = await res.json();
        setArtists((prev) =>
          prev.map((artist) =>
            artist.id === currentArtist.id ? updatedArtist : artist,
          ),
        );
      } else {
        res = await fetch("http://localhost:8080/api/admin/artists", {
          method: "POST",
          credentials: "include",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(payload),
        });

        if (!res.ok) throw new Error("Failed to create artist");
        updatedArtist = await res.json();

        setArtists((prev) => [...prev, updatedArtist]);
      }

      navigate("/admin/artists");
    } catch (err) {
      console.error("Error submitting artist:", err);
    }
  };

  const columns: { key: keyof Artist; header: string }[] = [
    { key: "id", header: "ID" },
    { key: "nickname", header: "Nickname" },
    { key: "bio", header: "Biography" },
    { key: "country", header: "Country" },
  ];

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-gray-800">Manage Artists</h2>
        <Button onClick={() => navigate("/admin/artists/new")}>
          <PlusIcon className="w-5 h-5 mr-1" />
          Добавить артиста
        </Button>
      </div>

      {isFormOpen ? (
        <div className="bg-white p-6 rounded-lg shadow-sm mb-6">
          <ArtistForm
            initialData={currentArtist}
            onSubmit={handleSubmit}
            onCancel={onCancel}
          />
        </div>
      ) : (
        <DataTable<Artist>
          columns={columns}
          data={artists}
          onEdit={handleEdit}
          onDelete={handleDelete}
        />
      )}
    </div>
  );
};

export default AdminArtistsPage;
