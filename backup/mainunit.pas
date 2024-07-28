unit MainUnit;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, Menus, ComCtrls,
  StdCtrls, ExtCtrls, Buttons, Grids, AnchorDockPanel;

type

  { TMainForm }

  TMainForm = class(TForm)
    Button1: TButton;
    MainMenu: TMainMenu;
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

{ TMainForm }

end.
