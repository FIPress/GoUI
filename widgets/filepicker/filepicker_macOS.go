// +build darwin,amd64,!sim

package filepicker

/*

#cgo darwin CFLAGS: -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa -framework WebKit

#include <Cocoa/Cocoa.h>
#include "../../c/common.h"

extern void filePicked(const char* s);

//title, nameLabel, message, canCreateDir, showsHidden, can

typedef struct FilePickerSettings{
    const char* title;
	const char* message;
    const char* fileTypes;
    const char* startLocation;
    const char* suggestedFilename;
    int multiple;
    int allowsFile;
    int allowsDir;
    int allowsOtherFileTypes;
    int canCreateDir;
    int showsHiddenFiles;
} FilePickerSettings;

void setProperties(NSSavePanel* panel, FilePickerSettings s ) {
	[panel setAllowedFileTypes:[[NSString stringWithUTF8String:s.fileTypes] componentsSeparatedByString:@";"]];
	[panel setAllowsOtherFileTypes:s.allowsOtherFileTypes];
	if(notEmpty(s.message)) {
		[panel setMessage:[NSString stringWithUTF8String:s.message]];
	}
	if(notEmpty(s.title)) {
		[panel setTitle:[NSString stringWithUTF8String:s.title]];
	}

	if(notEmpty(s.startLocation))  {
		[panel setDirectoryURL:[NSURL URLWithString:[NSString stringWithUTF8String:s.startLocation]]];
	}

	[panel setCanCreateDirectories:s.canCreateDir];
	[panel setShowsHiddenFiles:s.showsHiddenFiles];
}

char* openFilePicker(FilePickerSettings s) {
	NSOpenPanel *panel = [NSOpenPanel openPanel];
	[panel setAllowsMultipleSelection:s.multiple];
	[panel setCanChooseDirectories:s.allowsDir];
	[panel setCanChooseFiles:s.allowsFile];

	setProperties(panel,s);

	if ([panel runModal] == NSModalResponseOK) {
		int count = panel.URLs.count;
		if(count == 1) {
			NSString *path = [panel.URLs.firstObject path];
			return (char*)[path UTF8String];
		} else {
			NSMutableString *paths = [NSMutableString stringWithString: [panel.URLs.firstObject path]];
			for(int i=1;i<panel.URLs.count;i++) {
				[paths appendFormat: @";%@",[panel.URLs[i] path] ];
			}
			return (char*)[paths UTF8String];
		}
	}

	return "";
}

char* saveFilePicker(FilePickerSettings s) {
	NSSavePanel* savePanel = [NSSavePanel savePanel];
	setProperties(savePanel,s);
	NSModalResponse r =[savePanel runModal];
    if ( r== NSModalResponseOK)
    {
        NSURL *URL = [savePanel URL];
        if (URL)
        {
            NSString *path = [URL path];
            return (char*)[path UTF8String];
        }
    }
	return "";
}
*/
import "C"

func toCocoaSettings(settings *FilePickerSettings) C.FilePickerSettings {
	return C.FilePickerSettings{C.CString(settings.Title),
		C.CString(settings.Message),
		C.CString(settings.Accept),
		C.CString(settings.StartLocation),
		C.CString(settings.SuggestedFileName),
		BoolToCInt(settings.Multiple),
		BoolToCInt(!settings.DirOnly),
		BoolToCInt(!settings.FileOnly),
		BoolToCInt(settings.AllowsOtherFileTypes),
		BoolToCInt(settings.CanCreateDir),
		BoolToCInt(settings.ShowsHiddenFiles),
	}
}

func openFilePicker(settings *FilePickerSettings) string {
	//var cPath C.CString
	//var cPath string
	s := toCocoaSettings(settings)
	if settings.IsSave {
		return C.GoString(C.saveFilePicker(s))
	} else {
		return C.GoString(C.openFilePicker(s))
	}
	//return string(cPath)
}
