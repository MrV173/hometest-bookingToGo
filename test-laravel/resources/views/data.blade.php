<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>test-BookingToGo</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
</head>
<body>
    <div class="container fw-bold w-50 p-3">
        <span class="fw-bold fs-3">USER</span>
        <form action="" method="post">
            @csrf 
        <div class="mb-3 mt-3">
            <label for="nama" class="form-label">Nama</label>
            <input type="text" class="form-control" id="nama" placeholder="Masukan nama anda" name="nama" required>
        </div>
        <div class="mb-3">
            <label for="tanggal_lahir" class="form-label">Tanggal Lahir</label>
            <input type="date" class="form-control" id="tanggal_lahir" placeholder="Pilih tanggal" name="tanggal_lahir" required>
        </div>
        <div class="mb-3">
            <label for="email" class="form-label">Email</label>
            <input type="email" class="form-control" id="email" placeholder="Masukan email anda" name="email" required>
        </div>
        <div class="mb-3">
            <label for="telepon" class="form-label">No. Telephone</label>
            <input type="text" class="form-control" id="telepon" placeholder="Masukan nomor handphone anda" name="telepon" required>
        </div>
        <div class="mb-3">
            <label for="nationality" class="form-label">Kewarganegaraan</label>
            <select name="nationality_id" id="nationality" class="form-control">
                <option value="" hidden>Pilih Kewarganegaraan</option>
                @foreach ($nationalities as $nationality)
                    <option value="{{ $nationality['id'] }}">{{ $nationality['kewarganegaraan'] }}</option>
                @endforeach
            </select>
        </div>
        <hr />
        <div class="mt-3 mb-3 row"> 
            <div class="col">
                <span>Keluarga</span>
            </div>
            <div class="col">
                <a href="" class="text-decoration-underline"> +Tambah keluarga </a>
            </div>
        </div>
        <div class="mb-3 row">
            <div class="col">
                <label for="hubungan" class="form-label">Hubungan</label>
                <input type="text" class="form-control" id="hubungan" placeholder="masukan hubungan keluarga" name="hubungan" required>
            </div>
            <div class="col">
                <label for="rl_nama" class="form-label" >Nama</label>
                <input type="text" class="form-control" id="rl_nama" placeholder="Masukan Nama" name="rl_nama" required>
            </div>
            <div class="col">
                <label for="rl_tanggal_lahir" class="form-label">Tanggal lahir</label>
                <input type="date" class="form-control" id="rl_tanggal_lahir" placeholder="Pilih tanggal" name="rl_tanggal_lahir" required>
            </div>
            <div class="col pt-2">
                <button type="button" class="btn btn-danger bg-danger mt-4">Hapus</button>
            </div>
        </div>
        <div class="mb-3 row">
            <div class="col">
                <label for="hubungan2" class="form-label">Hubungan</label>
                <input type="text" class="form-control" id="hubungan2" placeholder="masukan hubungan keluarga" name="hubungan2" required>
            </div>
            <div class="col">
                <label for="rl_nama2" class="form-label" >Nama</label>
                <input type="text" class="form-control" id="rl_nama2" placeholder="Masukan Nama" name="rl_nama2" required>
            </div>
            <div class="col">
                <label for="rl_tanggal_lahir2" class="form-label">Tanggal lahir</label>
                <input type="date" class="form-control" id="rl_tanggal_lahir2" placeholder="Pilih tanggal" name="rl_tanggal_lahir2" required>
            </div>
            <div class="col pt-2">
                <button type="button" class="btn btn-danger bg-danger mt-4">Hapus</button>
            </div>
        </div>
        <div>
            <input type="submit" class="btn btn-primary" name="inputData">
        </div>
        </form>
    </div>
</body>
</html>