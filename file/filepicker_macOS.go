// +build darwin,amd64,!sim

package file

/*

#cgo darwin CFLAGS: -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa -framework WebKit

//title, nameLabel, message, canCreateDir, showsHidden, can

char* openFilePicker(bool multiple, bool dirOnly, char* defaultDir) {
	NSOpenPanel *panel = [NSOpenPanel openPanel];
	[panel directoryURL:NSHomeDirectory()];
 	[panel setAllowsMultipleSelection:NO];
	[panel setCanChooseDirectories:YES];
	[panel setCanChooseFiles:YES];
	[panel setAllowedFileTypes:@[@"onecodego"]];
    [panel setAllowsOtherFileTypes:YES];
    if ([panel runModal] == NSOKButton) {
        NSString *path = [panel.URLs.firstObject path];
        return [path UTF8String];
    }
	return "";
}

char* saveFilePicker() {
	NSSavePanel*    panel = [NSSavePanel savePanel];
	[panel setNameFieldStringValue:@"Untitle.onecodego"];
    [panel setMessage:@"Choose the path to save the document"];
	[panel setAllowsOtherFileTypes:YES];
	[panel setAllowedFileTypes:@[@"onecodego"]];
	[panel setExtensionHidden:YES];
	[panel setCanCreateDirectories:YES];
	[panel beginSheetModalForWindow:self.window completionHandler:^(NSInteger result){
		if (result == NSFileHandlingPanelOKButton)
		{
			NSString *path = [[panel URL] path];
       		[@"onecodego" writeToFile:path atomically:YES encoding:NSUTF8StringEncoding error:nil];
		}
	}];
}
*/
import "C"
