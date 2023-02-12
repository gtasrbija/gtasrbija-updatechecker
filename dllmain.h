
// +build cg

#include <windows.h>

void MainThread();


/* DWORD WINAPI MyThreadFunction() {
    OnProcessAttach();
    return 0;
} */

BOOL APIENTRY DllMain(HMODULE hModule, DWORD reason, LPVOID lpReserved) {
	if (reason == DLL_PROCESS_ATTACH) {
		DisableThreadLibraryCalls(hModule);
		CreateThread(NULL, 0, MainThread, 0, 0, NULL);
	}

	return TRUE;
}
