# Name of the installer
Outfile "agent_smith_installer.exe"

# Installation directory
InstallDir $PROGRAMFILES\Agent-Smith

# Request admin privileges for installation (necessary for Program Files)
RequestExecutionLevel admin

# Default section
Section "Install Agent-Smith"

    # Set the installation directory
    SetOutPath $INSTDIR

    # Copy the executable to the installation directory
    File "agent-smith.exe"

    # Write an uninstaller
    WriteUninstaller "$INSTDIR\uninstall.exe"

    # Add a registry entry to run the app at startup
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Run" "AgentSmith" "$INSTDIR\agent-smith.exe"
SectionEnd

# Uninstaller Section
Section "Uninstall"
    # Remove files and shortcuts
    Delete "$INSTDIR\agent-smith.exe"

    # Remove uninstaller
    Delete "$INSTDIR\uninstall.exe"

    # Remove the installation directory
    RMDir "$INSTDIR"

SectionEnd

