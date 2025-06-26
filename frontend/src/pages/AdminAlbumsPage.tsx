import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import DataTable, { type Column } from "../components/blocks/admin/DataTable";
import AlbumForm from "../components/form/admin/AlbumForm";
import { Button } from "../components/ui/Button";
import { PlusIcon } from "../components/ui/icons";
import type { Album } from "../types";

const AdminAlbumsPage = () => {
  const [albums, setAlbums] = useState<Album[]>([]);
  const [currentAlbum, setCurrentAlbum] = useState<Album | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();
  const { id } = useParams();
  const isFormOpen = !!id;

  const onCancel = () => {
    navigate("/admin/albums");
  };

  useEffect(() => {
    const fetchAlbums = async () => {
      setIsLoading(true);
      try {
        const res = await fetch("http://localhost:8080/api/albums", {
          method: "GET",
          credentials: "include",
        });
        const data = await res.json();
        if (Array.isArray(data)) {
          setAlbums(data);
        }
      } catch (error) {
        console.error("Error fetching albums:", error);
      } finally {
        setIsLoading(false);
      }
    };
    fetchAlbums();
  }, []);

  useEffect(() => {
    if (id && id !== "new") {
      const numericId = Number(id);
      const fetchAlbum = async () => {
        try {
          const res = await fetch(
            `http://localhost:8080/api/albums/${numericId}`,
            { credentials: "include" },
          );
          const data = await res.json();
          setCurrentAlbum(data);
        } catch (error) {
          console.error("Error fetching album:", error);
        }
      };
      fetchAlbum();
    } else {
      setCurrentAlbum(null);
    }
  }, [id]);

  const handleEdit = (album: Album) => {
    navigate(`/admin/albums/${album.id}`);
  };

  const handleDelete = async (id: number) => {
    try {
      const res = await fetch(`http://localhost:8080/api/admin/albums/${id}`, {
        method: "DELETE",
        credentials: "include",
      });
      if (!res.ok) throw new Error("Failed to delete album");
      setAlbums((prev) => prev.filter((album) => album.id !== id));
    } catch (err) {
      console.error("Error deleting album:", err);
    }
  };

  const handleSubmit = async (albumData: Album, imageFile?: File) => {
    const formData = new FormData();

    // Для PATCH запроса отправляем только измененные поля
    if (currentAlbum) {
      if (!imageFile && currentAlbum.image_url) {
        formData.append("current_image", currentAlbum.image_url);
      }

      if (albumData.title !== currentAlbum.title) {
        formData.append("title", albumData.title);
      }
      if (imageFile) {
        formData.append("image", imageFile);
      }
      if (albumData.artist.id !== currentAlbum.artist.id) {
        formData.append("artist_id", albumData.artist.id.toString());
      }
      // Для POST запроса отправляем все поля
    } else {
      formData.append("title", albumData.title);
      if (imageFile) {
        formData.append("image", imageFile);
      }
      formData.append("artist_id", albumData.artist.id.toString());
    }

    try {
      let res: Response;
      let url: string;
      let method: string;

      if (currentAlbum) {
        url = `http://localhost:8080/api/admin/albums/${currentAlbum.id}`;
        method = "PATCH";
      } else {
        url = `http://localhost:8080/api/admin/albums`;
        method = "POST";
      }

      // eslint-disable-next-line prefer-const
      res = await fetch(url, {
        method,
        credentials: "include",
        body: formData,
      });

      if (!res.ok) {
        const errorData = await res.json();
        throw new Error(errorData.message || "Request failed");
      }

      const updatedAlbum = await res.json();

      if (currentAlbum) {
        setAlbums((prev) =>
          prev.map((album) =>
            album.id === currentAlbum.id ? updatedAlbum : album,
          ),
        );
      } else {
        const res = await fetch("http://localhost:8080/api/albums", {
          method: "GET",
          credentials: "include",
        });
        const data = await res.json();
        if (Array.isArray(data)) {
          setAlbums(data);
        }
      }

      navigate("/admin/albums");
    } catch (err) {
      console.error("Error submitting album:", err);
    }
  };

  const columns: Column<Album>[] = [
    { key: "id", header: "ID" },
    { key: "title", header: "Title" },
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
          alt="Album"
          className="h-10 w-10 object-cover rounded"
        />
      ),
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
        }
        return "N/A";
      },
    },
  ];

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-gray-800">Manage Albums</h2>
        <Button onClick={() => navigate("/admin/albums/new")}>
          <PlusIcon className="w-5 h-5 mr-1" />
          Добавить альбом
        </Button>
      </div>

      {isFormOpen ? (
        <div className="bg-white p-6 rounded-lg shadow-sm mb-6">
          <AlbumForm
            initialData={currentAlbum}
            onSubmit={handleSubmit}
            onCancel={onCancel}
          />
        </div>
      ) : (
        <DataTable<Album>
          columns={columns}
          data={albums}
          onEdit={handleEdit}
          onDelete={handleDelete}
        />
      )}
    </div>
  );
};

export default AdminAlbumsPage;
