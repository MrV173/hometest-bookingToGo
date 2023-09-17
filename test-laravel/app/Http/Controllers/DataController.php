<?php

namespace App\Http\Controllers;
use Illuminate\Support\Facades\Http;


use Illuminate\Http\Request;

class DataController extends Controller
{
    public function index() {
        $responseNationalities = Http::get('http://localhost:5000/api/v1/nationalities');
        $nationailities = $responseNationalities['data'];
        return view('data', ['nationalities' => $nationailities]);
    }

    public function inputData(Request $request) {
        {
            // Data User/Customer yang akan di input
            $nama = $request->input('nama');
            $tanggal_lahir = $request->input('tanggal_lahir');
            $telepon = $request->input("telepon");
            $nationalityid = $request->input('nationality_id');
            $email = $request->input('email');

            $nationalityInt = intval($nationalityid);

            $data = [
                'nama' => $nama,
                'tanggal_lahir' => $tanggal_lahir,
                'telepon' => $telepon,
                'nationality_id' => $nationalityInt,
                'email' => $email,
            ];
            // Kirim permintaan POST ke Golang 
            $response = Http::post('http://localhost:5000/api/v1/customer', $data);

            if ($response->successful()) {

                $newResp = $response->json();
                // mengambil data id dari response untuk relasi ke table family list
                $ID = $newResp['data']['id'];

                // Data Keluarga yang akan di input
                $hubungan = $request->input('hubungan');
                $rl_nama = $request->input('rl_nama');
                $rl_tanggal_lahir = $request->input('rl_tanggal_lahir');
                $customer_id = $ID;
                
                $dataRL = [
                    'hubungan' => $hubungan,
                    'rl_nama' => $rl_nama,
                    'rl_tanggal_lahir' => $rl_tanggal_lahir,
                    'customer_id' => $customer_id,
                ];

                //Data keluarga ke 2 yang akan di input

                $hubungan2 = $request->input('hubungan2');
                $rl_nama2 = $request->input('rl_nama2');
                $rl_tanggal_lahir2 = $request->input('rl_tanggal_lahir2');
                $customer_id = $ID;
                
                $dataRL2 = [
                    'hubungan' => $hubungan2,
                    'rl_nama' => $rl_nama2,
                    'rl_tanggal_lahir' => $rl_tanggal_lahir2,
                    'customer_id' => $customer_id,
                ];

                // Melakukan POST pada data keluarga
                $responseData1 = Http::post('http://localhost:5000/api/v1/family', $dataRL);
                // Melakukan POST pada data keluarga2
                $responseData2 = Http::post('http://localhost:5000/api/v1/family', $dataRL2);
                // mengambil/ melakukan output data response dan melakukan GET ke Golang berdasrkan ID
                $output = Http::get("http://localhost:5000/api/v1/customer/$ID");
                dd($output->json(), "Data Berhasil Ditambahkan");

            } else {
                // Tangani respons tidak berhasil dari Golang (jika perlu)
                dd($response->json());
                return response()->json(['message' => 'Failed to create user'], $response->status());
            }
        }
    }
}
